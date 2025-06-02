package main

import "fmt"

type TransformInt func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2, 3}
	doubles := transformNumbers(&numbers, double)
	fmt.Println("Doubled:", doubles)
	triples := transformNumbers(&numbers, triple)
	fmt.Println("Tripled:", triples)
	transformed1 := transformNumbers(&numbers, getTransformerFn(&numbers))
	fmt.Println("Transformed 1:", transformed1)
	transformed2 := transformNumbers(&moreNumbers, getTransformerFn(&moreNumbers))
	fmt.Println("Transformed 2:", transformed2)
}

func transformNumbers(numbers *[]int, transform TransformInt) []int {
	dNumbers := make([]int, len(*numbers))
	for ind, val := range *numbers {
		dNumbers[ind] = transform(val)
	}
	return dNumbers
}

func getTransformerFn(numbers *[]int) TransformInt {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
