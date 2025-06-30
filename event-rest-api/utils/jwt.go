package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "0fe83316-ebe0-45e1-bf50-2b3b7d1690c0"

// The GenerateToken function creates a JWT token with user email and a placeholder user ID.
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(4 * time.Hour), // Token expires in 4 hours
	})

	return token.SignedString([]byte("secretKey"))
}
