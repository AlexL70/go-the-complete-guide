package main

import (
	"errors"
	"fmt"

	"example.com/note/note" // Adjust the import path as necessary
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}
	fmt.Println(note)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Enter the title of the note: ")
	if err != nil {
		return "", "", err
	}
	content, err := getUserInput("Enter the content of the note: ")
	if err != nil {
		return "", "", err
	}
	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	if value == "" {
		return "", errors.New("Input cannot be empty")
	}
	return value, nil
}
