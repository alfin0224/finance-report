package routes

import (
	"finance-report/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/products", controllers.GetProducts)
		v1.POST("/products", controllers.CreateProduct)
		// Define other CRUD routes for products
	}
	{
		v1.GET("/finance_reports", controllers.GetFinanceReports)
	}
	// Define routes for categories and finance reports as well
}
