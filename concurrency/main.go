package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
}

func main() {
	go greet("Nice to meet you!")
	go greet("How are you?")
	go slowGreet("How ... are ... you ... doing?")
	go greet("I hope you are liking the course!")
}
