package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")

	var appUser User = User{
		FirstName: userFirstName,
		LastName:  userLastName,
		BirthDate: userBirthDate,
		CreatedAt: time.Now(),
	}

	printUserData(&appUser)
}

func printUserData(user *User) {
	fmt.Printf("User: %s %s, Birth Date: %s, Created At: %s\n",
		// Pointers to structs allow accessing fields with or without dereferencing
		(*user).FirstName, user.LastName, user.BirthDate, user.CreatedAt.Format(time.RFC3339))
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
