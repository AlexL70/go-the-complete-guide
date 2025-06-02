package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	doubled := transformNumbers(&numbers, func(i int) int {
		return i * 2
	})
	fmt.Println(doubled)
	tripled := transformNumbers(&numbers, createTransformer(3))
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	transformedNumbers := make([]int, len(*numbers))
	for i, num := range *numbers {
		transformedNumbers[i] = transform(num)
	}
	return transformedNumbers
}

// createTransformer returns a function that multiplies an integer by a given factor.
func createTransformer(factor int) func(int) int {
	return func(i int) int {
		return i * factor // using the factor to transform the number is called closure
	}
}
