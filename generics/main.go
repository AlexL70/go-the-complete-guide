package main

import "fmt"

func main() {
	result := add(3, 4)
	fmt.Println("Sum of integers:", result)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
