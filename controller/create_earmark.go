package controller

import (
	"net/http"

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

	resp.ResponseCode = "0000"
	resp.ResponseMessage = "Earmark Processed successfully"
	resp.ResponseStatus = "S"

	// Send JSON response
	c.JSON(http.StatusOK, resp)
}
