package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/db"
	"rest-api/routes"
)

func main() {
	db.InitDb()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
