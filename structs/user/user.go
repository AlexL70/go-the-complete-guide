package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate time.Time // Using time.Time for date representation
	createdAt time.Time
}

// Constructors are not a common pattern in Go, but we can use methods to initialize structs.
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("First name and last name cannot be empty")
	}
	var date, err = time.Parse("01/02/2006", birthDate)
	if err != nil {
		return nil, fmt.Errorf("please enter date in the next format \"MM/DD/YYYY\"; error parsing date: %v", err)
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: date,
		createdAt: time.Now(),
	}, nil
}

func (u User) PrintUserData() {
	fmt.Printf("User: %s %s, Birth Date: %s, Created At: %s\n",
		u.firstName, u.lastName, u.birthDate.Format(time.DateOnly), u.createdAt.Format(time.RFC3339))
}

// To update structure fields, we must use a pointer receiver
// otherwise we will modify a copy of the struct and the original will remain unchanged.
func (u *User) ClearUserName() {
	// Pointers to structs allow accessing fields with or without dereferencing
	(*u).firstName = ""
	u.lastName = ""
}
