package main

import (
	controllers "go-dummy-app/controller"
	mongoConnection "go-dummy-app/database"
	producer "go-dummy-app/kafka"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoConnection.ConnectMongo()

	producer.Produce()

	r := gin.Default()

	// Route to get dummy user
	r.GET("/user", controllers.CreateAndStoreDummyUser)
	r.GET("/users", controllers.GetAllUsers)

	r.Run() // default port is :8080
}
