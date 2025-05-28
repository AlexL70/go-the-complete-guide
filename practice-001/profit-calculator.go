package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	tax := ebt * (taxRate / 100)
	profit := ebt - tax
	fmt.Printf("Profit: %.2f\n", profit)
	ratio := ebt / profit
	fmt.Printf("EBT to Profit Ratio: %.2f\n", ratio)
}
