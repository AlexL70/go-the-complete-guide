package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.0 // Annual inflation rate in percentage

func main() {
	var investmentAmount, expectedReturnRate, years float64

	outputText("Enter the amount you want to invest: $")
	fmt.Scan(&investmentAmount)
	outputText("Enter the expected annual return rate (in percentage): ")
	fmt.Scan(&expectedReturnRate)
	outputText("Enter the number of years you want to invest for: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calcFutureValues(investmentAmount, expectedReturnRate, years)
	output := fmt.Sprintf("Future Value of Investment: $%.2f\nFuture Value adjusted to inflation rate: $%.2f\n", futureValue, futureRealValue)
	fmt.Print(output)
}

func outputText(text string) {
	fmt.Print(text)
}

func calcFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
