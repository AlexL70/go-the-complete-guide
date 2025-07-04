package routes

import (
	"event-rest-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("error parsing event ID: %s", err)})
		return
	} else if eventId <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "event ID must be a positive integer"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("Event with ID %d not found: %w", eventId, err)})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error registering for event: %s", err)})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("error parsing event ID: %s", err)})
		return
	} else if eventId <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "event ID must be a positive integer"})
		return
	}

	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error canceling registration: %s", err)})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully canceled registration for event"})
}
