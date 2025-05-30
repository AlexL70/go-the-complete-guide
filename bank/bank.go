package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, fmt.Errorf("Error reading file: %v", err)
	}
	var value float64
	value, err = strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, fmt.Errorf("Error parsing float from file: %v", err)
	}
	return value, nil
}

func writeFloatToFile(value float64, fileName string) {
	balanceStr := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(balanceStr), 0644)
}

func main() {
	var accountBalance float64 = 1000.00 // Initial balance
	var savedBalance, err = getFloatFromFile(accountBalanceFile)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		fmt.Println("No previous balance found, starting with default balance of $1000.00")
	} else {
		accountBalance = savedBalance
		fmt.Printf("Previous balance loaded: $%.2f\n", accountBalance)
	}

	fmt.Println("Welcome to the Bank Management System")
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
					writeFloatToFile(accountBalance, accountBalanceFile)
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
					writeFloatToFile(accountBalance, accountBalanceFile)
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
