package routes

import (
	"github.com/gin-gonic/gin"
	"go-sandbox.uz/course-project/middlewares"
)

// By pointer we are using original gin.Engine instance, not a copy
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)

}

// as the function in events.go in the same package, we can call it directly without any import, and in lowercase
