package photos

import "github.com/labstack/echo/v4"

type Photo struct {
	AlbumId       int    `json:"albumId"`
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Url           string `json:"url"`
	ThumnbnailUrl string `json:"thumnbnailUrl"`
}

func PhotosGroupHandler(photo *echo.Group) {
	photo.GET("/", GetPhotos)
	photo.GET("/album/:albumid", GetPhotosByAlbumId)
	photo.GET("/id/:id", GetPhotoById)
}
