package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate time.Time // Using time.Time for date representation
	CreatedAt time.Time
}

// Constructors are not a common pattern in Go, but we can use methods to initialize structs.
func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("First name and last name cannot be empty")
	}
	var date, err = time.Parse("01/02/2006", birthDate)
	if err != nil {
		return nil, fmt.Errorf("please enter date in the next format \"MM/DD/YYYY\"; error parsing date: %v", err)
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: date,
		CreatedAt: time.Now(),
	}, nil
}

func (u User) printUserData() {
	fmt.Printf("User: %s %s, Birth Date: %s, Created At: %s\n",
		u.FirstName, u.LastName, u.BirthDate.Format("DD of MMM, YYYY"), u.CreatedAt.Format(time.RFC3339))
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

	var appUser *User
	appUser, err := NewUser(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	appUser.printUserData()
	appUser.clearUserName()
	appUser.printUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
