package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	//server.POST("/events", middlewares.Authenticate, createEvent) // gin supports chaining functions, they executed in order. We can register middleware for authentication. But better is to use Grouped() with Use().
	//server.PUT("/events/:id", updateEvent)
	//server.DELETE("/events/:id", deleteEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", middlewares.Authenticate, createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
