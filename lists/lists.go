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
	featuredPrices[0] = 30.99 // Changing the first element of the slice
	fmt.Println("Updated featured prices:", featuredPrices)
	fmt.Println("Updated prices:", prices)
	toTheEnd := prices[1:] // Slicing from index 2 to the end
	fmt.Println("Prices from index 1 to the end:", toTheEnd)
	fromTheBeginning := prices[:3] // Slicing from the beginning to index 2 (exclusive)
	fmt.Println("Prices from the beginning to index 2:", fromTheBeginning)

	fmt.Println("Length and cap of featuredPrices:", len(featuredPrices), cap(featuredPrices))
	fmt.Println("Length and cap of prices:", len(prices), cap(prices))
	fmt.Println("Length and cap of from the beginning prices:", len(fromTheBeginning), cap(fromTheBeginning))
	fromTheBeginning = fromTheBeginning[:4] // Extending the slice to include all elements again
	fmt.Println("Updated length and cap of from the beginning prices:", len(fromTheBeginning), cap(fromTheBeginning))
	fmt.Println("Updated from the beginning prices:", fromTheBeginning)
}
