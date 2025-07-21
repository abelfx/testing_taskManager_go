package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection
var UserCollection *mongo.Collection

func InitDB() {
	clientOptions := options.Client().ApplyURI(("mongodb://localhost:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("Mongo connection error: ", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("Mongo ping error:", err)
	}

	TaskCollection = client.Database("taskdb").Collection("tasks")
	UserCollection = client.Database("userdb").Collection("users")
	log.Println("Connected to MongoDB")

}