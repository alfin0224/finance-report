package routes

import (
	"finance-report/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/products", controllers.CreateProduct)
		// Define other CRUD routes for products
	}
	// Define routes for categories and finance reports as well
}
