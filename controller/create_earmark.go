package controller

import (
	"net/http"

	"context"
	"go-dummy-app/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	models "go-dummy-app/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func CreateEarmark(c *gin.Context) {
	var req models.EarmarkRequest

	// Simulate creating a user (in real case, save to DB)
	var resp models.EarmarkResponse

	// Validate & bind JSON input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		resp.ResponseCode = "0024"
		resp.ResponseMessage = "Earmark Validation Failed"
		resp.ResponseStatus = "F"
		return
	}

	copier.Copy(&resp, &req)

	existingEarmarks := fetchExistingEarmark(c)

	if req.RequestType == "CREATE" {
		if existingEarmarks != nil {
			resp.ResponseCode = "0001"
			resp.ResponseMessage = "Earmark Already exists"
			resp.ResponseStatus = "F"
		} else {
			var earmarkStatus models.EarmarkStatus
			copier.Copy(&earmarkStatus, &req)
			earmarkStatus.Status = "ACTIVE"

			var earmarkStatusCollection *mongo.Collection = database.Client.Database("testdb").Collection("earmark_status")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			_, err := earmarkStatusCollection.InsertOne(ctx, earmarkStatus)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
				return
			}

			resp.ResponseCode = "0000"
			resp.ResponseMessage = "Earmark Processed successfully"
			resp.ResponseStatus = "S"
		}
	} else {
		existingEarmarks[0].Status = "CLOSED"
		resp.ResponseCode = "0008"
		resp.ResponseMessage = "Earmark Released successfully"
		resp.ResponseStatus = "S"
	}

	// Send JSON response
	c.JSON(http.StatusOK, resp)
}

func fetchExistingEarmark(c *gin.Context) []models.EarmarkStatus {
	// Get the collection
	earmarkStatusCollection := database.Client.Database("testdb").Collection("earmark_status")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB query to get all documents
	cursor, err := earmarkStatusCollection.Find(ctx, bson.M{})
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
