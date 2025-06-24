package main

import (
	"event-rest-api/db"
	"event-rest-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	var server = gin.Default()
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events", createEvent)
	server.Run(":8080") // Start the server on port 8080 (on the local machine)
}

func getEventById(context *gin.Context) {
	id_str := context.Param("id")
	id, err := strconv.ParseInt(id_str, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid ID: %s. You must pass an integer value.", id_str)})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
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
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusCreated, event)
}
