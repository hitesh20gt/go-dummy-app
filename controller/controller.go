package controller

import (
	"context"
	"go-dummy-app/database"
	"go-dummy-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"

	kafkaConnection "go-dummy-app/kafka"
)

func CreateAndStoreDummyUser(c *gin.Context) {
	dummyUser := models.User{
		ID:    uuid.New().String(),
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	var userCollection *mongo.Collection = database.Client.Database("testdb").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, dummyUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}

	kafkaConnection.Produce()

	c.JSON(http.StatusOK, dummyUser)
}
