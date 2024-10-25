package routes

import (
	"go_rest_api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// Event routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	
	// Group routes that require authentication
	authenticated := server.Group("/")

	// Use the Authenticate middleware for all routes in the authenticated group
	authenticated.Use(middlewares.Authenticate)

	//Event handlers
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	//Event handlers
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistrationForEvent) 

	// User routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}