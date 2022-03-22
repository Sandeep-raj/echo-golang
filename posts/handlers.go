package posts

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
	var postData []Post
	err := json.Unmarshal(data, &postData)
	if err != nil {
		return err
	}

	wrtModel := []mongo.WriteModel{}

	for _, post := range postData {
		updateModel := mongo.NewUpdateOneModel().
			SetFilter(bson.D{primitive.E{Key: "id", Value: post.Id}}).
			SetUpdate(bson.D{primitive.E{Key: "$set", Value: post}}).SetUpsert(true)
		wrtModel = append(wrtModel, updateModel)
	}

	result, err := mongoclient.PostCol.BulkWrite(context.Background(), wrtModel)

	if err != nil {
		return err
	}

	log.Printf("Sucessfully inserted the record. [%+v]", result)

	return nil
}

func GetPosts(c echo.Context) error {
	var result []Post
	cursor, err := mongoclient.PostCol.Find(context.Background(), bson.D{})
	log.Print(cursor)
	if err != nil {
		return err
	}
	cursor.All(context.Background(), &result)
	return c.JSON(http.StatusOK, result)
}

func GetPostsByUserId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Err while converting the id to interger. err [%+v]", err))
	}

	resultList := []Post{}

	cursor, err := mongoclient.PostCol.Find(context.Background(), bson.D{primitive.E{Key: "userid", Value: id}})
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Err while getting data from post collection. err [%+v]", err))
	}

	err = cursor.All(context.Background(), &resultList)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Err while parsing the data from collection. err [%+v]", err))
	}

	if len(resultList) == 0 {
		return c.String(http.StatusNotFound, fmt.Sprintf("err while fetching data for %d id for posts", id))
	}

	return c.JSON(http.StatusOK, resultList)
}

func GetPostsById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Err while converting the id to interger. err [%+v]", err))
	}

	var result Post

	curr := mongoclient.PostCol.FindOne(context.Background(), bson.D{primitive.E{Key: "id", Value: id}})
	curr.Decode(&result)

	if result.UserId == 0 {
		return c.String(http.StatusNotFound, fmt.Sprintf("err while fetching data for %d id for posts", id))
	}

	return c.JSON(http.StatusOK, result)
}
