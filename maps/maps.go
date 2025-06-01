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
	fmt.Println("Websites:", websites)
}
