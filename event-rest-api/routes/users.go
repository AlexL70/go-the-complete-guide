package routes

import (
	"event-rest-api/models"
	"fmt"
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
