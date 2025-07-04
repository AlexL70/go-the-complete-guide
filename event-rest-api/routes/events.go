package routes

import (
	"event-rest-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	event.UserID = context.GetInt64("userId")
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusCreated, event)
}

func updateEvent(context *gin.Context) {
	id_str := context.Param("id")
	id, err := strconv.ParseInt(id_str, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid ID: %s. You must pass an integer value.", id_str)})
		return
	}
	originalEvent, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	if originalEvent.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this event, because you are not the owner."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}

func deleteEvent(context *gin.Context) {
	id_str := context.Param("id")
	id, err := strconv.ParseInt(id_str, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid ID: %s. You must pass an integer value.", id_str)})
		return
	}
	originalEvent, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	if originalEvent.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this event, because you are not the owner."})
		return
	}

	err = originalEvent.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
