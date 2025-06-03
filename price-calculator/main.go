package main

import (
	"fmt"

	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		job := prices.NewTaxIncludedPriceJob(taxRate)
		job.Process()
		fmt.Printf("Tax Rate: %.2f\n", taxRate)
		fmt.Println("Price with taxes:", job.TaxIncludedPrices)
	}
}
