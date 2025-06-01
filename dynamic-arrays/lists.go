package main

import "fmt"

func main() {
	prices := []float64{10.99, 12.49}
	fmt.Println("Slice only first element:", prices[0:1])
	// prices[2] = 15.99 // this will cause a runtime error instead we could use append
	updatedPrices := append(prices, 15.99, 18.49, 20.00, 22.50)
	fmt.Println("Updated prices:", updatedPrices)
	fmt.Println("Original prices:", prices)
	diminishingPrices := prices[1:]
	fmt.Println("Diminishing prices:", diminishingPrices)
	discountedPrices := []float64{10.99, 12.49, 15.99, 18.49, 20.00, 22.50}
	mergedPrices := append(updatedPrices, discountedPrices...)
	fmt.Println("Merged prices:", mergedPrices)
}
