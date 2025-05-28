package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.0 // Annual inflation rate in percentage
	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Enter the amount you want to invest: $")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the expected annual return rate (in percentage): ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter the number of years you want to invest for: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Printf("Future Value of Investment: $%.2f\n", futureValue)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Printf("Future Value adjusted to inflation rate: $%.2f\n", futureRealValue)
}
