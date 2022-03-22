package todos

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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMockData(data []byte) error {
	var TodoData []Todo
	err := json.Unmarshal(data, &TodoData)
	if err != nil {
		return err
	}

	wrtModel := []mongo.WriteModel{}

	for _, Todo := range TodoData {
		updateModel := mongo.NewUpdateOneModel().
			SetFilter(bson.D{primitive.E{Key: "id", Value: Todo.Id}}).
			SetUpdate(bson.D{primitive.E{Key: "$set", Value: Todo}}).SetUpsert(true)
		wrtModel = append(wrtModel, updateModel)
	}

	result, err := mongoclient.TodoCol.BulkWrite(context.Background(), wrtModel)

	if err != nil {
		return err
	}

	log.Printf("Sucessfully inserted the record. [%+v]", result)

	return nil
}

func GetTodos(c echo.Context) error {
	var result []Todo
	cursor, err := mongoclient.TodoCol.
		Find(context.Background(),
			bson.D{},
			options.Find().SetProjection(bson.D{primitive.E{Key: "_id", Value: 0}}))

	if err != nil {
		return err
	}
	cursor.All(context.Background(), &result)
	return c.JSON(http.StatusOK, result)
}

func GetTodosByUserId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Err while converting the id to interger. err [%+v]", err))
	}

	resultList := []Todo{}

	cursor, err := mongoclient.TodoCol.Find(context.Background(), bson.D{primitive.E{Key: "userid", Value: id}})
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Err while getting data from Todos collection. err [%+v]", err))
	}

	err = cursor.All(context.Background(), &resultList)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Err while parsing the data from collection. err [%+v]", err))
	}

	if len(resultList) == 0 {
		return c.String(http.StatusNotFound, fmt.Sprintf("err while fetching data for %d id for Todos", id))
	}

	return c.JSON(http.StatusOK, resultList)
}

func GetTodosById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Err while converting the id to interger. err [%+v]", err))
	}

	var result Todo

	curr := mongoclient.TodoCol.FindOne(context.Background(), bson.D{primitive.E{Key: "id", Value: id}})
	curr.Decode(&result)

	if result.UserId == 0 {
		return c.String(http.StatusNotFound, fmt.Sprintf("err while fetching data for %d id for Todos", id))
	}

	return c.JSON(http.StatusOK, result)
}
