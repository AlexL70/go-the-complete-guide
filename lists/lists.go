package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	prices := [4]float64{10.99, 20.49, 5.99, 15.00}
	fmt.Println("Prices:", prices)
	fmt.Println("Product names:", productNames)
	fmt.Println("Third price:", prices[2])
	// Slicing the array
	featuredPrices := prices[1:3] // Slicing from index 1 to 2 (exclusive)
	fmt.Println("Featured prices:", featuredPrices)
}
