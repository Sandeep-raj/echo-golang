package users

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/yellowmessenger/mongoclient"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitMockData(data []byte) error {
	var UserData []User
	err := json.Unmarshal(data, &UserData)
	if err != nil {
		return err
	}

	wrtModel := []mongo.WriteModel{}

	for _, User := range UserData {
		updateModel := mongo.NewUpdateOneModel().
			SetFilter(bson.D{primitive.E{Key: "id", Value: User.Id}}).
			SetUpdate(bson.D{primitive.E{Key: "$set", Value: User}}).SetUpsert(true)
		wrtModel = append(wrtModel, updateModel)
	}

	result, err := mongoclient.UserCol.BulkWrite(context.Background(), wrtModel)

	if err != nil {
		return err
	}

	log.Printf("Sucessfully inserted the record. [%+v]", result)

	return nil
}

func GetUsers(c echo.Context) error {
	var result []User
	cursor, err := mongoclient.UserCol.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	cursor.All(context.Background(), &result)
	return c.JSON(http.StatusOK, result)
}

func GetUsersByUsername(c echo.Context) error {
	username := c.Param("username")
	if username == "" {
		return c.String(http.StatusBadRequest, "Err while getting username.")
	}

	resultList := []User{}

	cursor, err := mongoclient.UserCol.Find(context.Background(), bson.D{primitive.E{Key: "username", Value: username}})
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Err while getting data from Users collection. err [%+v]", err))
	}

	err = cursor.All(context.Background(), &resultList)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Err while parsing the data from collection. err [%+v]", err))
	}

	if len(resultList) == 0 {
		return c.String(http.StatusNotFound, fmt.Sprintf("err while fetching data for %s id for Users", username))
	}

	return c.JSON(http.StatusOK, resultList)
}

func GetUsersById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Err while converting the id to interger. err [%+v]", err))
	}

	var result User

	curr := mongoclient.UserCol.FindOne(context.Background(), bson.D{primitive.E{Key: "id", Value: id}})
	curr.Decode(&result)

	if result.Username == "" {
		return c.String(http.StatusNotFound, fmt.Sprintf("err while fetching data for %d id for Users", id))
	}

	return c.JSON(http.StatusOK, result)
}
