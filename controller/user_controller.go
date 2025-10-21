package controller

import (
	"context"
	"go-clean-app/database"
	"go-clean-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUsers(c *gin.Context) {
	// Get the collection
	userCollection := database.Client.Database("testdb").Collection("users")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB query to get all documents
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer cursor.Close(ctx)

	var users []models.User

	// Iterate over the cursor
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding user"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
