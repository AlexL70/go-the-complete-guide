package main

import (
	"fmt"

	"example.com/price-calculator/prices"
)

func main() {
	pricesList := []float64{10.0, 20.0, 30.0}
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		job := prices.NewTaxIncludedPriceJob(taxRate, pricesList)
		job.Process()
		fmt.Printf("Tax Rate: %.2f\n", taxRate)
		fmt.Println("Price with taxes:", job.TaxIncludedPrices)
	}
}
