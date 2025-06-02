package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	transformed := transformNumbers(&numbers, func(i int) int {
		return i * 2
	})
	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	transformedNumbers := make([]int, len(*numbers))
	for i, num := range *numbers {
		transformedNumbers[i] = transform(num)
	}
	return transformedNumbers
}
