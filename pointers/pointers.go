package main

import "fmt"

func main() {
	age := 32          // regular variable
	agePointer := &age // pointer to the variable

	fmt.Printf("Age: %d\n", *agePointer)
	fmt.Printf("Age Pointer: %p\n", agePointer)
	editAgeToAdultYears(agePointer)      // pass pointer to function
	fmt.Printf("Adult Years: %d\n", age) // print updated age
}

func editAgeToAdultYears(age *int) {
	*age -= 18 // calculate adult years and update the value
}
