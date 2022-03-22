package albums

import "github.com/labstack/echo/v4"

type Album struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}

func AlbumGroupHandler(album *echo.Group) {
	album.GET("/", GetAlbums)
	album.GET("/userid/:userid", GetAlbumsByUserId)
	album.GET("/id/:id", GetAlbumById)
}
