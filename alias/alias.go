package main

import "fmt"

type customString string

func (s customString) log() {
	fmt.Println(s)
}

func main() {
	var str customString = "Hello, World!"
	str.log()
}
