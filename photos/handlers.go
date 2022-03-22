package photos

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
	var photoData []Photo
	err := json.Unmarshal(data, &photoData)
	if err != nil {
		return err
	}

	wrtModel := []mongo.WriteModel{}

	for _, photo := range photoData {
		updateModel := mongo.NewUpdateOneModel().
			SetFilter(bson.D{primitive.E{Key: "id", Value: photo.Id}}).
			SetUpdate(bson.D{primitive.E{Key: "$set", Value: photo}}).SetUpsert(true)
		wrtModel = append(wrtModel, updateModel)
	}

	result, err := mongoclient.PhotoCol.BulkWrite(context.Background(), wrtModel)

	if err != nil {
		return err
	}

	log.Printf("Sucessfully inserted the record. [%+v]", result)

	return nil
}

func GetPhotos(c echo.Context) error {
	var result []Photo
	cursor, err := mongoclient.PhotoCol.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	cursor.All(context.Background(), &result)
	return c.JSON(http.StatusOK, result)
}

func GetPhotosByAlbumId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("albumid"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Err while converting the id to interger. err [%+v]", err))
	}

	resultList := []Photo{}

	cursor, err := mongoclient.PhotoCol.Find(context.Background(), bson.D{primitive.E{Key: "albumid", Value: id}})
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Err while getting data from Photos collection. err [%+v]", err))
	}

	err = cursor.All(context.Background(), &resultList)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Err while parsing the data from collection. err [%+v]", err))
	}

	if len(resultList) == 0 {
		return c.String(http.StatusNotFound, fmt.Sprintf("err while fetching data for %d id for Photos", id))
	}

	return c.JSON(http.StatusOK, resultList)
}

func GetPhotoById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Err while converting the id to interger. err [%+v]", err))
	}

	var result Photo

	curr := mongoclient.PhotoCol.FindOne(context.Background(), bson.D{primitive.E{Key: "id", Value: id}})
	curr.Decode(&result)

	if result.AlbumId == 0 {
		return c.String(http.StatusNotFound, fmt.Sprintf("err while fetching data for %d id for Photos", id))
	}

	return c.JSON(http.StatusOK, result)
}
