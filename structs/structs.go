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

func (u User) printUserData() {
	fmt.Printf("User: %s %s, Birth Date: %s, Created At: %s\n",
		u.FirstName, u.LastName, u.BirthDate, u.CreatedAt.Format(time.RFC3339))
}

// To update structure fields, we must use a pointer receiver
// otherwise we will modify a copy of the struct and the original will remain unchanged.
func (u *User) clearUserName() {
	// Pointers to structs allow accessing fields with or without dereferencing
	(*u).FirstName = ""
	u.LastName = ""
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

	appUser.printUserData()
	appUser.clearUserName()
	appUser.printUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
