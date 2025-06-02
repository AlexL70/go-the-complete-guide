package main

import "fmt"

func main() {
	numbers := []int{1, 12, 23, 34, 45}
	sum := sumUp(numbers[0], numbers[1:]...)
	fmt.Println("The sum is:", sum)
	sumVar := sumUp(12, 23, 34, 45, 56)
	fmt.Println("The sum of variable arguments is:", sumVar)
}

func sumUp(theFirst int, numbers ...int) int {
	sum := theFirst
	for _, number := range numbers {
		sum += number
	}
	return sum
}
