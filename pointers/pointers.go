package main

import "fmt"

func main() {
	age := 32          // regular variable
	agePointer := &age // pointer to the variable

	fmt.Printf("Age: %d\n", *agePointer)
	fmt.Printf("Age Pointer: %p\n", agePointer)
	fmt.Printf("Adult Years: %d\n", getAdultYears(age))
}

func getAdultYears(age int) int {
	return age - 18 // calculate adult years
}
