package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.00 // Initial balance

	fmt.Println("Welcome to the Bank Management System")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Println("Please enter your choice (1-4):")
	fmt.Scanln(&choice)
	if choice < 1 || choice > 4 {
		fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		return
	}
	switch choice {
	case 1:
		fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
	default:
		fmt.Println("Thank you for using the Bank Management System. Goodbye!")
	}
}
