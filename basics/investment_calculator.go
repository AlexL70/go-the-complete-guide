package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.0 // Annual inflation rate in percentage
	var investmentAmount, expectedReturnRate, years float64

	outputText("Enter the amount you want to invest: $")
	fmt.Scan(&investmentAmount)
	outputText("Enter the expected annual return rate (in percentage): ")
	fmt.Scan(&expectedReturnRate)
	outputText("Enter the number of years you want to invest for: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	output := fmt.Sprintf("Future Value of Investment: $%.2f\nFuture Value adjusted to inflation rate: $%.2f\n", futureValue, futureRealValue)
	fmt.Print(output)
}

func outputText(text string) {
	fmt.Print(text)
}
