package main

import "fmt"

func main() {
	var investmentAmount = 1000.0
	var expectedReturnRate = 5.5
	// var years = 10

	var futureValue = investmentAmount * (1 + expectedReturnRate/100)
	fmt.Println("Future Value of Investment: $", futureValue)
}
