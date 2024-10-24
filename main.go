package main

import (
	"go_rest_api/db"
	"go_rest_api/routes"

	"github.com/gin-gonic/gin"
)


func main()  {

	// Initialize the database
	db.InitDB()

	// Create a new gin router
	server := gin.Default()

	// Register the routes
	routes.RegisterRoutes(server)

	// Run the server
	server.Run(":8080") // localhost:8080
}

