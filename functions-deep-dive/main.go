package main

import "fmt"

type TransformInt func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubles := transformNumbers(&numbers, double)
	fmt.Println("Doubled:", doubles)
	triples := transformNumbers(&numbers, triple)
	fmt.Println("Tripled:", triples)
}

func transformNumbers(numbers *[]int, transform TransformInt) []int {
	dNumbers := make([]int, len(*numbers))
	for ind, val := range *numbers {
		dNumbers[ind] = transform(val)
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
