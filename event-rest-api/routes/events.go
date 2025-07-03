package routes

import (
	"event-rest-api/models"
	"event-rest-api/utils"
	"fmt"
	"log"
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
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	_, err := utils.ValidateToken(token)
	if err != nil {
		log.Printf("Error validating token: %w", err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	// email := (*claims)["email"].(string)
	// userId := (*claims)["userId"].(int64)
	var userId = int64(1) // Placeholder for user ID, replace with actual user ID extraction logic

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.UserID = userId
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
	_, err = models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
