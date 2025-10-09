package handlers

import (
	"net/http"

	"github.com/atilio70/stock-control-sistem/backend/internal/database"
	"github.com/atilio70/stock-control-sistem/backend/internal/models"
	"github.com/gin-gonic/gin"
)

// GetClients - List all clients
func GetClients(c *gin.Context) {
	var clients []models.Client
	if err := database.DB.Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

// GetClientByID - Get a single client by ID
func GetClientsByID(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := database.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	c.JSON(http.StatusOK, client)
}

// CreateClient - Add a new client
func CreateClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, client)
}

// UpdateClient -Edit an existing client
func UpdateClient(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := database.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&client)
	c.JSON(http.StatusOK, client)
}

// DeleteClient - Remove a client
func DeleteClient(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := database.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	database.DB.Delete(&client)
	c.JSON(http.StatusOK, gin.H{"message": "Client deleted"})
}
