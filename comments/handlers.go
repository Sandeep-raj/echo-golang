package comments

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
	var commentData []Comment
	err := json.Unmarshal(data, &commentData)
	if err != nil {
		return err
	}

	wrtModel := []mongo.WriteModel{}

	for _, comment := range commentData {
		updateModel := mongo.NewUpdateOneModel().
			SetFilter(bson.D{primitive.E{Key: "id", Value: comment.Id}}).
			SetUpdate(bson.D{primitive.E{Key: "$set", Value: comment}}).SetUpsert(true)
		wrtModel = append(wrtModel, updateModel)
	}

	result, err := mongoclient.CommentCol.BulkWrite(context.Background(), wrtModel)

	if err != nil {
		return err
	}

	log.Printf("Sucessfully inserted the record. [%+v]", result)

	return nil
}

func GetComments(c echo.Context) error {
	var result []Comment
	cursor, err := mongoclient.CommentCol.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	cursor.All(context.Background(), &result)
	return c.JSON(http.StatusOK, result)
}

func GetCommentsByPostId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("postid"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Err while converting the id to interger. err [%+v]", err))
	}

	resultList := []Comment{}

	cursor, err := mongoclient.CommentCol.Find(context.Background(), bson.D{primitive.E{Key: "postid", Value: id}})
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Err while getting data from comments collection. err [%+v]", err))
	}

	err = cursor.All(context.Background(), &resultList)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Err while parsing the data from collection. err [%+v]", err))
	}

	if len(resultList) == 0 {
		return c.String(http.StatusNotFound, fmt.Sprintf("err while fetching data for %d id for comments", id))
	}

	return c.JSON(http.StatusOK, resultList)
}

func GetCommentsById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Err while converting the id to interger. err [%+v]", err))
	}

	var result Comment

	curr := mongoclient.CommentCol.FindOne(context.Background(), bson.D{primitive.E{Key: "id", Value: id}})
	curr.Decode(&result)

	if result.PostId == 0 {
		return c.String(http.StatusNotFound, fmt.Sprintf("err while fetching data for %d id for comments", id))
	}

	return c.JSON(http.StatusOK, result)
}
