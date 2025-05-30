package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	revenue, expenses, taxRate, err := inputValues()
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	ebt, profit, ratio := calcEbtTaxProfitAndRatio(revenue, expenses, taxRate)
	var result = fmt.Sprintf("Earnings Before Tax (EBT): %.2f\n", ebt)
	result += fmt.Sprintf("Profit: %.2f\n", profit)
	result += fmt.Sprintf("EBT to Profit Ratio: %.2f\n", ratio)
	fmt.Println(result)
	os.WriteFile("profit_report.txt", []byte(result), 0644)
}

func inputValues() (revenue, expenses, taxRate float64, err error) {
	revenue, err = getInput("Revenue: ")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid input for revenue: %v", err)
	}
	expenses, err = getInput("Expenses: ")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid input for expenses: %v", err)
	}
	taxRate, err = getInput("Tax Rate: ")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid input for tax rate: %v", err)
	}
	return revenue, expenses, taxRate, nil
}

func getInput(prompt string) (float64, error) {
	var value float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	_, err = fmt.Sscanf(input, "%f", &value)
	if err != nil {
		return 0, fmt.Errorf("input must be a positive number")
	}
	if value <= 0 {
		return 0, fmt.Errorf("input must be a positive number")
	}
	return value, nil
}

func calcEbtTaxProfitAndRatio(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	tax := ebt * (taxRate / 100)
	profit = ebt - tax
	ratio = ebt / profit
	return
}
