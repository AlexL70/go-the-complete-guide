package middlewares

import (
	"event-rest-api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	claims, err := utils.ValidateToken(token)
	if err != nil {
		log.Printf("Error validating token: %v", err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	email := (*claims)["email"].(string)
	userId := (*claims)["userId"].(float64)

	context.Set("userId", int64(userId))
	context.Set("email", email)
	context.Next()
}
