package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/alfin0224/finance-report/controllers"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/products", controllers.CreateProduct)
		// Implement other routes here
	}

	return r
}
