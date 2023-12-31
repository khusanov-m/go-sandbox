package routes

import "github.com/gin-gonic/gin"

// By pointer we are using original gin.Engine instance, not a copy
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", addEvent)
  server.PUT("/events/:id", updateEvent)
}

// as the function in events.go in the same package, we can call it directly without any import, and in lowercase
