package main

import "fmt"

func main() {
	age := 32 // regular variable
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Adult Years: %d\n", getAdultYears(age))
}

func getAdultYears(age int) int {
	return age - 18 // calculate adult years
}
