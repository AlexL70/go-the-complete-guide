package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	prices := [4]float64{10.99, 20.49, 5.99, 15.00}
	fmt.Println("Prices:", prices)
	fmt.Println("Product names:", productNames)
	fmt.Println(prices[2])
}
