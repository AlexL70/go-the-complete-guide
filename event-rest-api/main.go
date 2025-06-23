package main

import (
	"event-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var server = gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080") // Start the server on port 8080 (on the local machine)
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, event)
}
