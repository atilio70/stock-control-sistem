package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts is a placeholder handler for retrieving products
func GetProducts(c *gin.Context) {
	//Placeholder response
	c.JSON(http.StatusOK, gin.H{
		"products": []string{"Product 1", "Product 2"},
	})
}
