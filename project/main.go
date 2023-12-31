package main

import (
	"github.com/gin-gonic/gin"
	"go-sandbox.uz/course-project/db"
	"go-sandbox.uz/course-project/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
