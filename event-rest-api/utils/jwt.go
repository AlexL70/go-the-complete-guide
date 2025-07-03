package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "50458505-460f-43c1-bbe4-1a8bc610485a"

// The GenerateToken function creates a JWT token with user email and a placeholder user ID.
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().UTC().Add(time.Hour * 4).Unix(), // Token expires in 4 hours
	})

	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return &claims, nil
}
