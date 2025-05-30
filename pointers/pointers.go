package main

import "fmt"

func main() {
	age := 32          // regular variable
	agePointer := &age // pointer to the variable

	fmt.Printf("Age: %d\n", *agePointer)
	fmt.Printf("Age Pointer: %p\n", agePointer)
	getAdultYears(agePointer) // pass pointer to function
	fmt.Printf("Adult Years: %d\n", *agePointer)
}

func getAdultYears(age *int) {
	*age -= 18 // calculate adult years and update the value
}
