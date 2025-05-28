package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.0 // Annual inflation rate in percentage
	var investmentAmount float64 = 1000
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future Value of Investment: $", futureValue)
	fmt.Println("Future Value adjusted to inflation rate: $", futureRealValue)
}
