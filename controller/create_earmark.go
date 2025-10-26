package controller

import (
	"fmt"
	"net/http"

	"context"
	"go-dummy-app/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	models "go-dummy-app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	existingEarmarks := fetchExistingEarmark(c, req.EarmarkReference)

	if req.RequestType == "CREATE" {
		if existingEarmarks != nil {
			resp.ResponseCode = "0001"
			resp.ResponseMessage = "Earmark Already exists"
			resp.ResponseStatus = "F"
		} else {
			var earmarkStatus models.EarmarkStatus
			copier.Copy(&earmarkStatus, &req)
			earmarkStatus.Status = "ACTIVE"
			earmarkStatus.ID = uuid.NewString()

			var earmarkStatusCollection *mongo.Collection = database.Client.Database("testdb").Collection("earmark_status")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			_, err := earmarkStatusCollection.InsertOne(ctx, earmarkStatus)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
				resp.ResponseCode = "0007"
				resp.ResponseMessage = "Intenal server error"
				resp.ResponseStatus = "F"
				return
			}

			resp.ResponseCode = "0000"
			resp.ResponseMessage = "Earmark Processed successfully"
			resp.ResponseStatus = "S"
		}
	} else {
		found, ok := Find(existingEarmarks, func(n models.EarmarkStatus) bool {
			return n.Status == "ACTIVE"
		})
		if ok {
			fmt.Println("Found:", found)
			found.Status = "CLOSED"
			var earmarkStatusCollection *mongo.Collection = database.Client.Database("testdb").Collection("earmark_status")

			filter := bson.M{"_id": found.ID}
			update := bson.M{"$set": found}

			opts := options.Update().SetUpsert(true)

			result, err := earmarkStatusCollection.UpdateOne(context.TODO(), filter, update, opts)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
				resp.ResponseCode = "0007"
				resp.ResponseMessage = "Intenal server error"
				resp.ResponseStatus = "F"
				return
			}
			if result.UpsertedCount > 0 {
				fmt.Println("✅ New document inserted with ID:", result.UpsertedID)
			} else {
				fmt.Println("✅ Existing document updated")
			}
			resp.ResponseCode = "0008"
			resp.ResponseMessage = "Earmark Released successfully"
			resp.ResponseStatus = "S"
		} else {
			fmt.Println("No match found")
			resp.ResponseCode = "0001"
			resp.ResponseMessage = "Earmark doesnt exists"
			resp.ResponseStatus = "F"
		}
	}

	// Send JSON response
	c.JSON(http.StatusOK, resp)
}

func fetchExistingEarmark(c *gin.Context, ermrkRef string) []models.EarmarkStatus {
	// Get the collection
	earmarkStatusCollection := database.Client.Database("testdb").Collection("earmark_status")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB query to get all documents
	filter := bson.M{"earmarkReference": ermrkRef}
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
func Find[T any](arr []T, condition func(T) bool) (T, bool) {
	for _, v := range arr {
		if condition(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}
