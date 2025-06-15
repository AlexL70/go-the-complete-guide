package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var server = gin.Default()
	server.GET("/events", getEvents)
	server.Run(":8080") // Start the server on port 8080 (on the local machine)
}

func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello! Eventually, this endpoint will return a list of events",
	})
}
