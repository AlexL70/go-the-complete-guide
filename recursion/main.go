package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println("Factorial of 5 is:", fact)
}

// factorial of 5 would be 5*4*3*2*1
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
