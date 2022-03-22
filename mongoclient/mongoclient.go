package mongoclient

import (
	"context"
	"log"
	"time"

	"bitbucket.org/yellowmessenger/configmanager"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var PostCol *mongo.Collection
var CommentCol *mongo.Collection
var AlbumCol *mongo.Collection
var PhotoCol *mongo.Collection
var TodoCol *mongo.Collection
var UserCol *mongo.Collection

func InitMongo() error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(configmanager.ConfStore.MongoConnStr))
	if err != nil {
		log.Printf("Error connecting mongo. Err [%+v]", err.Error())
		return err
	}

	for _, col := range configmanager.ConfStore.DatasetList {
		switch col {
		case "posts":
			PostCol = Client.Database(configmanager.ConfStore.Database).Collection("posts")
		case "comments":
			CommentCol = Client.Database(configmanager.ConfStore.Database).Collection("comments")
		case "albums":
			AlbumCol = Client.Database(configmanager.ConfStore.Database).Collection("albums")
		case "photos":
			PhotoCol = Client.Database(configmanager.ConfStore.Database).Collection("photos")
		case "todos":
			TodoCol = Client.Database(configmanager.ConfStore.Database).Collection("todos")
		case "users":
			UserCol = Client.Database(configmanager.ConfStore.Database).Collection("users")
		default:
			log.Printf("No collection with %s is found.", col)
		}
	}

	return err
}

func GetClient(colname string) *mongo.Collection {
	return Client.Database(configmanager.ConfStore.Database).Collection(colname)
}
