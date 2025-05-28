package main

import "fmt"

func main() {
	revenue, expenses, taxRate := inputValues()

	ebt, profit, ratio := calcEbtTaxProfitAndRatio(revenue, expenses, taxRate)
	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("EBT to Profit Ratio: %.2f\n", ratio)
}

func inputValues() (revenue, expenses, taxRate float64) {
	revenue = getInput("Revenue: ")
	expenses = getInput("Expenses: ")
	taxRate = getInput("Tax Rate: ")
	return
}

func getInput(prompt string) float64 {
	var value float64
	fmt.Print(prompt)
	fmt.Scan(&value)
	return value
}

func calcEbtTaxProfitAndRatio(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	tax := ebt * (taxRate / 100)
	profit = ebt - tax
	ratio = ebt / profit
	return
}
