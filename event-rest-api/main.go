package main

import (
	"event-rest-api/db"
	"event-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	var server = gin.Default()
	routes.RegisterRoutes(server) // Register the routes defined in the routes package
	server.Run(":8080")           // Start the server on port 8080 (on the local machine)
}
