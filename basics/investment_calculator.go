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
	fmt.Println("Future Value of Investment: $", futureValue)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future Value adjusted to inflation rate: $", futureRealValue)
}
