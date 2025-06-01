package main

import "fmt"

func main() {
	// websites := []string{
	// 	"https://www.google.com",
	// 	"https://aws.com"}
	// fmt.Println("Websites:", websites)
	websites := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://aws.com",
		"GitHub":              "https://github.com",
	}
	// Print the map
	fmt.Println("Websites:", websites)
	// Accessing a specific value
	fmt.Println("AWS URL:", websites["Amazon Web Services"])
	// Adding a new key-value pair
	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println("Updated Websites:", websites)
	// removing an existing key-value pair
	delete(websites, "GitHub")
	fmt.Println("After deletion:", websites)
}
