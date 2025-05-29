package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.00 // Initial balance

	fmt.Println("Welcome to the Bank Management System")
MAIN_LOOP:
	for {
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
			continue
		} else {
			switch choice {
			case 1:
				fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
			case 2:
				var depositAmount float64
				fmt.Println("Enter the amount to deposit:")
				fmt.Scanln(&depositAmount)
				if depositAmount <= 0 {
					fmt.Println("Deposit amount must be greater than zero.")
				} else {
					accountBalance += depositAmount
					fmt.Printf("You have successfully deposited $%.2f. New balance is: $%.2f\n", depositAmount, accountBalance)
				}
			case 3:
				var withdrawAmount float64
				fmt.Println("Enter the amount to withdraw:")
				fmt.Scanln(&withdrawAmount)
				if withdrawAmount <= 0 {
					fmt.Println("Withdrawal amount must be greater than zero.")
				} else if withdrawAmount > accountBalance {
					fmt.Println("Insufficient funds for this withdrawal.")
				} else {
					accountBalance -= withdrawAmount
					fmt.Printf("You have successfully withdrawn $%.2f. New balance is: $%.2f\n", withdrawAmount, accountBalance)
				}
			default:
				fmt.Println("Exiting bank system.")
				break MAIN_LOOP
			}
		}
	}
	fmt.Println("Thank you for choosing our bank management system! Goodbye!")
}
