// product_controller.go
package controllers

import (
	"finance-report/api/models"
	"finance-report/config"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	db, err := config.SetupDatabase()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to connect to database"})
		return
	}

	var products []models.Product
	result := db.Find(&products)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(200, products)
}

func CreateProduct(c *gin.Context) {
	db, err := config.SetupDatabase()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to connect to database"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	result := db.Create(&product)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(201, product)
}
