package main

import "fmt"

type floatMap map[string]float64

func (f floatMap) output() {
	for key, value := range f {
		fmt.Printf("%s: %.2f\n", key, value)
	}
}

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
	fmt.Println("User names after appending additional names:")
	for i, name := range userNames {
		fmt.Printf("%d: %s\n", i+1, name)
	}
	fmt.Println("Total user names:", len(userNames))
	// Still have te original capacity because enough space was allocated originally
	// No new allocation happened
	fmt.Println("Capacity of user names slice:", cap(userNames))
	fmt.Println()

	// Allocating maps
	courseRatings := make(floatMap, 5) // Preallocate space for 5 course ratings
	courseRatings["Go Programming"] = 4.5
	courseRatings["Python Programming"] = 4.7
	courseRatings["Java Programming"] = 4.2
	courseRatings["JavaScript Programming"] = 4.8
	// Go does not allocate new space yet for keys because 5 is the initial capacity
	courseRatings["C++ Programming"] = 4.3
	// And here the new space is allocated because the initial capacity is exceeded
	courseRatings["Ruby Programming"] = 4.6
	fmt.Println("Course ratings:")
	courseRatings.output()
	fmt.Println("Total course ratings:", len(courseRatings))
}
