package main

import "fmt"

func main() {
	userNames := make([]string, 6, 20)
	// Fill in existing user names
	userNames[0] = "Max"
	userNames[1] = "Manuel"
	userNames[2] = "Pedro"
	userNames[3] = "John"
	userNames[4] = "Paul"
	userNames[5] = "George"
	// userNames[6] = "Ringo" // This will cause a runtime error because the slice has a length of 6
	fmt.Println("User names:", userNames)
	additionalNames := []string{"Alice", "Bob", "Charlie"}
	// Append additional user names
	userNames = append(userNames, additionalNames...)
	fmt.Println("User names after appending additional names:", userNames)
	fmt.Println("User names:", userNames)
	fmt.Println("Total user names:", len(userNames))
	// Still have te original capacity because enough space was allocated originally
	// No new allocation happened
	fmt.Println("Capacity of user names slice:", cap(userNames))
}
