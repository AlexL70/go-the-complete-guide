package main

import "fmt"

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data
	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
