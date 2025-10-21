package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}

	// Ping to verify connection
	if err = Client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo ping error:", err)
	}

	log.Println("Connected to MongoDB")
}
