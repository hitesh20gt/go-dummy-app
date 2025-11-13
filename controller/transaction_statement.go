package controller

import (
	"net/http"

	"context"
	"go-dummy-app/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	models "go-dummy-app/models"

	"github.com/gin-gonic/gin"
)

func TransactionStatement(c *gin.Context) {
	var req models.EarmarkStatusRequest

	// Simulate creating a user (in real case, save to DB)
	var existingEarmarks []models.EarmarkStatus

	// Validate & bind JSON input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	existingEarmarks = fetchExistingEarmarkStatus(c, req.DebitAccount)

	// Send JSON response
	c.JSON(http.StatusOK, existingEarmarks)
}

func fetchExistingEarmarkStatus(c *gin.Context, account string) []models.EarmarkStatus {
	// Get the collection
	earmarkStatusCollection := database.Client.Database("testdb").Collection("earmark_status")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB query to get all documents
	filter := bson.M{"debitAccount": account}
	cursor, err := earmarkStatusCollection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return []models.EarmarkStatus{}
	}
	defer cursor.Close(ctx)

	var existingEarmarks []models.EarmarkStatus

	// Iterate over the cursor
	for cursor.Next(ctx) {
		var existingEarmark models.EarmarkStatus
		if err := cursor.Decode(&existingEarmark); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding user"})
			return []models.EarmarkStatus{}
		}
		existingEarmarks = append(existingEarmarks, existingEarmark)
	}

	return existingEarmarks
}

// Generic-style function (for Go 1.18+)
func FindEarmark[T any](arr []T, condition func(T) bool) (T, bool) {
	for _, v := range arr {
		if condition(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}
