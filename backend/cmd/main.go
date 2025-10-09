package main

import (
	"log"

	"github.com/atilio70/stock-control-sistem/backend/internal/database"
	"github.com/atilio70/stock-control-sistem/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

// main is the entry point of the application
func main() {
	//Initialize the database conection
	database.Connect()

	//create a Gin router
	router := gin.Default()

	//Products routes
	productRoutes := router.Group("/api/products")
	{
		productRoutes.GET("/", handlers.GetProducts)         //List all products
		productRoutes.GET("/:id", handlers.GetProductByID)   //Get a product by ID
		productRoutes.POST("/", handlers.CreateProduct)      //Create a product
		productRoutes.PUT("/:id", handlers.UpdateProduct)    //Update a product
		productRoutes.DELETE("/:id", handlers.DeleteProduct) //Delete a product
	}

	//Clients routes
	clientRoutes := router.Group("/api/clients")
	{
		clientRoutes.GET("/", handlers.GetClients)         //List all products
		clientRoutes.GET("/:id", handlers.GetClientsByID)  //Get a client by ID
		clientRoutes.POST("/", handlers.CreateClient)      //Create a new client
		clientRoutes.PUT("/:id", handlers.UpdateClient)    //Update client
		clientRoutes.DELETE("/:id", handlers.DeleteClient) //Delete a clients
	}

	//simple test endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Start the server on port 8080
	log.Println("Listening and serving HTTP on :8080")
	router.Run(":8080")
}
