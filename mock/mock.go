package mock

import (
	"log"

	"bitbucket.org/yellowmessenger/albums"
	"bitbucket.org/yellowmessenger/comments"
	"bitbucket.org/yellowmessenger/configmanager"
	"bitbucket.org/yellowmessenger/photos"
	"bitbucket.org/yellowmessenger/posts"
	restyClient "bitbucket.org/yellowmessenger/resty"
	"bitbucket.org/yellowmessenger/todos"
	"bitbucket.org/yellowmessenger/users"
)

func InitMock() {
	db_list := configmanager.ConfStore.DatasetList

	for _, db := range db_list {
		switch db {
		case "posts":
			InitPosts()
		case "comments":
			InitComments()
		case "albums":
			InitAlbums()
		case "photos":
			InitPhotos()
		case "todos":
			InitTodos()
		case "users":
			InitUsers()
		}
	}
}

func InitPosts() {
	resp, err := restyClient.Client.R().Get("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		log.Fatalf("Error while getting the posts data. Err [%+v]", err)
	}

	err = posts.InitMockData(resp.Body())
	if err != nil {
		log.Panicf("error while bootstrapping the post data. Err [%+v]", err)
	}
}

func InitComments() {
	resp, err := restyClient.Client.R().Get("https://jsonplaceholder.typicode.com/comments/")
	if err != nil {
		log.Fatalf("Error while getting the posts data. Err [%+v]", err)
	}

	err = comments.InitMockData(resp.Body())
	if err != nil {
		log.Panicf("error while bootstrapping the post data. Err [%+v]", err)
	}
}

func InitAlbums() {
	resp, err := restyClient.Client.R().Get("https://jsonplaceholder.typicode.com/albums/")
	if err != nil {
		log.Fatalf("Error while getting the posts data. Err [%+v]", err)
	}

	err = albums.InitMockData(resp.Body())
	if err != nil {
		log.Panicf("error while bootstrapping the post data. Err [%+v]", err)
	}
}

func InitPhotos() {
	resp, err := restyClient.Client.R().Get("https://jsonplaceholder.typicode.com/photos/")
	if err != nil {
		log.Fatalf("Error while getting the posts data. Err [%+v]", err)
	}

	err = photos.InitMockData(resp.Body())
	if err != nil {
		log.Panicf("error while bootstrapping the post data. Err [%+v]", err)
	}
}

func InitTodos() {
	resp, err := restyClient.Client.R().Get("https://jsonplaceholder.typicode.com/todos/")
	if err != nil {
		log.Fatalf("Error while getting the posts data. Err [%+v]", err)
	}

	err = todos.InitMockData(resp.Body())
	if err != nil {
		log.Panicf("error while bootstrapping the post data. Err [%+v]", err)
	}
}

func InitUsers() {
	resp, err := restyClient.Client.R().Get("https://jsonplaceholder.typicode.com/users/")
	if err != nil {
		log.Fatalf("Error while getting the posts data. Err [%+v]", err)
	}

	err = users.InitMockData(resp.Body())
	if err != nil {
		log.Panicf("error while bootstrapping the post data. Err [%+v]", err)
	}
}
