package main

imporr (
	"github.com/gin-gonic/gin"
	"github.com/atilio70/stock-control-sistem/backend/internal/database"
)

//main is the entry point of the application
func main() {
	//Initialize the database conection
	database.Connect()

	//create a Gin router
	router := gin.Default()

	//simple test route
	reoter.GET("/ping", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Start the server on port 8080
	router.Run(":8080")
}