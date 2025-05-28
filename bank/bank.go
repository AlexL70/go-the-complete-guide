package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Bank Management System")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Println("Please enter your choice (1-4):")
	fmt.Scanln(&choice)
	fmt.Println("You chose option:", choice)
}
