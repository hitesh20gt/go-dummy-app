package controller

import (
	"context"
	"go-clean-app/database"
	"go-clean-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateAndStoreDummyUser(c *gin.Context) {
	dummyUser := models.User{
		ID:    "1",
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

	c.JSON(http.StatusOK, dummyUser)
}
