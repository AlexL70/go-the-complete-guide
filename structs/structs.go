package main

import (
	"fmt"

	"example.com/structs/user" // Adjust the import path according to your project structure
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	appUser.PrintUserData()
	appUser.ClearUserName()
	appUser.PrintUserData()

	admin := user.NewAdmin("someemail@somesurl.com", "verySecureP@ssw0rd")
	admin.PrintUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
