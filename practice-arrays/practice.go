package main

import "fmt"

type Product struct {
	id    int
	title string
	price float64
}

func main() {
	// Time to practice what you learned!

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"Programming", "Hiking", "Playing Guitar"}
	fmt.Println("Hobbies:", hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println("First Hobby:", hobbies[0])
	hobbiesCombined := []string{hobbies[1], hobbies[2]}
	fmt.Println("Second and Third Hobbies Combined as List:", hobbiesCombined)
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	slice1 := hobbies[0:2] // First way: using slicing syntax
	slice2 := hobbies[:2]  // Second way: using shorthand syntax
	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	reSliced := slice1[1:3] // Re-slicing to get second and last element
	fmt.Println("Re-sliced Slice:", reSliced)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Learn Go", "Build a Web App"}
	fmt.Println("Goals:", goals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "Learn Web API with Go"
	goals = append(goals, "Pass Go tests on HackerRank")
	fmt.Println("Updated Goals:", goals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	prodList := []Product{
		{1, "Laptop", 999.99},
		{2, "Smartphone", 499.99},
	}
	fmt.Println("Product List:", prodList)
	prodList = append(prodList, Product{3, "Headphones", 199.99})
	fmt.Println("Expanded Product List:", prodList)
}
