package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Connect to //MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func GetDBDetails() *mongo.Client {
	return client
}

func GetDatabaseName() *mongo.Database {
	log.Println("Get Database Name")
	databaseName := client.Database("goTest")
	return databaseName
}

func FetchUsersCollection() *mongo.Collection {
	log.Println("Get User Collection Name")
	userCollection := client.Database("goTest").Collection("users")
	return userCollection
}

func FetchTodoCollection() *mongo.Collection {
	log.Println("Get Todo Collection Name")
	userCollection := client.Database("goTest").Collection("todo")
	return userCollection
}
