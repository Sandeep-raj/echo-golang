package main

import (
	"log"

	"bitbucket.org/yellowmessenger/albums"
	"bitbucket.org/yellowmessenger/comments"
	"bitbucket.org/yellowmessenger/configmanager"
	"bitbucket.org/yellowmessenger/mock"
	"bitbucket.org/yellowmessenger/mongoclient"
	"bitbucket.org/yellowmessenger/photos"
	"bitbucket.org/yellowmessenger/posts"
	restyClient "bitbucket.org/yellowmessenger/resty"
	"bitbucket.org/yellowmessenger/todos"
	"bitbucket.org/yellowmessenger/users"
	"github.com/labstack/echo/v4"
)

func init() {
	restyClient.InitRestyClient()
	err := configmanager.InitConfig()
	if err != nil {
		log.Fatalf("Error loading Configuration. Err [%+v]", err)
	}

	err = mongoclient.InitMongo()
	if err != nil {
		log.Fatalf("Error connecting Mongo. Err [%+v]", err)
	}

	mock.InitMock()
}

func main() {
	e := echo.New()

	posts.PostGroupHandler(e.Group("/posts"))
	comments.CommentGroupHandler(e.Group("/comments"))
	albums.AlbumGroupHandler(e.Group("/albums"))
	photos.PhotosGroupHandler(e.Group("/photos"))
	todos.TodosGroupHandler(e.Group("/todos"))
	users.UserGroupHandler(e.Group("/users"))

	e.Logger.Fatal(e.Start(":8088"))
}
