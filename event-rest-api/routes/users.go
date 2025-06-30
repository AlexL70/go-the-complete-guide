package routes

import (
	"event-rest-api/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Errorf("Could not parse request's data: %w", err),
		})
		return
	}

	user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("Could not save user in the DB: %w", err),
		})
		return
	}

	successMsg := fmt.Sprintf(`User created successfully: "%s"`, user.Email)
	context.JSON(http.StatusCreated, gin.H{
		"message": successMsg,
	})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Errorf("Could not parse request's data: %w", err),
		})
		return
	}
	err = models.ValidateCredentials(user.Email, user.Password)
	if err != nil {
		log.Println(fmt.Errorf("Invalid credentials for user with email %s: %w", user.Email, err))
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials or user not found or server error. Please try again. If the problem persists, contact support.",
		})
		return
	}
	log.Printf("User with email %s logged in successfully.\n", user.Email)
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User with email %s logged in successfully", user.Email),
	})
}
