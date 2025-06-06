package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance float64 = 1000.00 // Initial balance
	var savedBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		fmt.Println("No previous balance found, starting with default balance of $1000.00")
	} else {
		accountBalance = savedBalance
		fmt.Printf("Previous balance loaded: $%.2f\n", accountBalance)
	}

	fmt.Println("Welcome to the Bank Management System")
	fmt.Println("Our phone number is:", randomdata.PhoneNumber())
MAIN_LOOP:
	for {
		presentOptions()
		var choice int
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
					fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
					fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
