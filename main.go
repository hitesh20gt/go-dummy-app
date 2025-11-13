package main

import (
	controllers "go-dummy-app/controller"
	mongoConnection "go-dummy-app/database"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoConnection.ConnectMongo()

	// go kafkaConnection.Consume()

	r := gin.Default()

	// Route to get dummy user
	r.GET("/user", controllers.CreateAndStoreDummyUser)
	r.POST("/earmark", controllers.CreateEarmark)
	r.POST("/transactionStatement", controllers.TransactionStatement)
	r.GET("/users", controllers.GetAllUsers)

	r.Run() // default port is :8080
}
