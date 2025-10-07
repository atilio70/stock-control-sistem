package main

import (
	"github.com/atilio70/stock-control-sistem/backend/internal/database"
	"github.com/gin-gonic/gin"
)

// main is the entry point of the application
func main() {
	//Initialize the database conection
	database.Connect()

	//create a Gin router
	router := gin.Default()

	//simple test route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Start the server on port 8080
	router.Run(":8080")
}
