package routes

import (
	"event-rest-api/models"
	"event-rest-api/utils"
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

	err = user.Save()
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
	err = user.ValidateCredentials()
	if err != nil {
		log.Println(fmt.Errorf("Invalid credentials for user with email %s: %w", user.Email, err))
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials or user not found or server error. Please try again. If the problem persists, contact support.",
		})
		return
	}
	log.Printf("User with email %s logged in successfully.\n", user.Email)
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		log.Println(fmt.Errorf("Could not generate token for user with email %s: %w", user.Email, err))
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate token. Please try again later.",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User with email %s logged in successfully", user.Email),
		"token":   token,
	})
}
