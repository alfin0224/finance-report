// finance_report_controller.go
package controllers

import (
	"finance-report/api/models"
	"finance-report/config"

	"github.com/gin-gonic/gin"
)

func GetFinanceReports(c *gin.Context) {
	db, err := config.SetupDatabase()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to connect to database"})
		return
	}

	var financeReports []models.FinanceReport
	result := db.Find(&financeReports)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve finance reports"})
		return
	}

	c.JSON(200, financeReports)
}
