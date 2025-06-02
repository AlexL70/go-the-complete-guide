package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	transformed := transformNumbers(&numbers, ???) 
	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	transformedNumbers := make([]int, len(*numbers))
	for i, num := range *numbers {
		transformedNumbers[i] = transform(num)
	}
	return transformedNumbers
}
