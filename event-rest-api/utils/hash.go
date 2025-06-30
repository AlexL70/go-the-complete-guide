package utils

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), (bcrypt.DefaultCost+bcrypt.MaxCost)/2)
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println(fmt.Errorf("Error comparing passwords: %w", err))
		return false
	}
	return true
}
