package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {
	go greet("Nice to meet you!")
	go greet("How are you?")
	done := make(chan bool)
	go slowGreet("How ... are ... you ... doing?", done)
	go greet("I hope you are liking the course!")
	var finished bool
	finished = <-done // Wait for the slowGreet to finish
	fmt.Println("Finished slow greeting:", finished)
}
