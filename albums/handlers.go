package albums

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
	var albumData []Album
	err := json.Unmarshal(data, &albumData)
	if err != nil {
		return err
	}

	wrtModel := []mongo.WriteModel{}

	for _, album := range albumData {
		updateModel := mongo.NewUpdateOneModel().
			SetFilter(bson.D{primitive.E{Key: "id", Value: album.Id}}).
			SetUpdate(bson.D{primitive.E{Key: "$set", Value: album}}).SetUpsert(true)
		wrtModel = append(wrtModel, updateModel)
	}

	result, err := mongoclient.AlbumCol.BulkWrite(context.Background(), wrtModel)

	if err != nil {
		return err
	}

	log.Printf("Sucessfully inserted the record. [%+v]", result)

	return nil
}

func GetAlbums(c echo.Context) error {
	var result []Album
	cursor, err := mongoclient.AlbumCol.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	cursor.All(context.Background(), &result)
	return c.JSON(http.StatusOK, result)
}

func GetAlbumsByUserId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Err while converting the id to interger. err [%+v]", err))
	}

	resultList := []Album{}

	cursor, err := mongoclient.AlbumCol.Find(context.Background(), bson.D{primitive.E{Key: "userid", Value: id}})
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Err while getting data from albums collection. err [%+v]", err))
	}

	err = cursor.All(context.Background(), &resultList)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Err while parsing the data from collection. err [%+v]", err))
	}

	if len(resultList) == 0 {
		return c.String(http.StatusNotFound, fmt.Sprintf("err while fetching data for %d id for albums", id))
	}

	return c.JSON(http.StatusOK, resultList)
}

func GetAlbumById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Err while converting the id to interger. err [%+v]", err))
	}

	var result Album

	curr := mongoclient.AlbumCol.FindOne(context.Background(), bson.D{primitive.E{Key: "id", Value: id}})
	curr.Decode(&result)

	if result.UserId == 0 {
		return c.String(http.StatusNotFound, fmt.Sprintf("err while fetching data for %d id for albums", id))
	}

	return c.JSON(http.StatusOK, result)
}
