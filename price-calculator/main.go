package main

import "fmt"

func main() {
	prices := []float64{10.0, 20.0, 30.0}
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}
	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxIncludedPrices := make([]float64, 0, len(prices))
		for _, price := range prices {
			total := price * (1 + taxRate)
			taxIncludedPrices = append(taxIncludedPrices, total)
		}
		result[taxRate] = taxIncludedPrices
	}
	fmt.Println("Prices with Tax:", result)
}
