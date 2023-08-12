package main

import (
	"finance-report/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080") // Change the port as needed
}
