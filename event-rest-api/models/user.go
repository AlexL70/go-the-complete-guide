package models

import (
	"event-rest-api/db"
	"event-rest-api/utils"
	"fmt"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func ValidateCredentials(email, password string) error {
	var user User
	query := "SELECT id, email, password FROM users WHERE email = ?"
	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return fmt.Errorf("User with email %s not found: %w", email, err)
	}
	if !utils.ComparePasswords(user.Password, password) {
		return fmt.Errorf("Invalid credentials for user with email %s", email)
	}
	return nil
}
