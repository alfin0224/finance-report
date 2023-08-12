package main

import (
	"github.com/gin-gonic/gin"
	"github.com/alfin0224/go-crud-api/routes"
)

func main() {
	db, _ := ConnectDB()
	r := routes.SetupRouter(db)

	r.Run(":8080")
}
