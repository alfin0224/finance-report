package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/alfin0224/finance-report/models"
	"gorm.io/gorm"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// Implement Update, Delete, and Get handlers similarly
