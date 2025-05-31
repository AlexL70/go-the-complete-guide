package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type Saver interface {
	Save() error
}

type Outputtable interface {
	Display()
	Saver
}

func main() {
	printSomething("Welcome to the Note and Todo App!")
	printSomething(123.45)
	printSomething(true)
	printSomething(-123)
	printSomething([]string{"Hello", "World"})

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
	todoText, err := getUserInput("Enter a todo text: ")
	if err != nil {
		fmt.Println("Error reading todo data:", err)
		return
	}
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	output(note)
	output(todo)
}

func printSomething(something any) {
	switch v := something.(type) {
	case int:
		fmt.Printf("Printing integer: %d\n", v)
	case float64:
		fmt.Printf("Printing float: %.2f\n", v)
	case bool:
		fmt.Printf("Printing boolean: %t\n", v)
	default:
		fmt.Println(something)
	}
}

func output(data Outputtable) {
	data.Display()
	saveData(data)
}

func saveData(data Saver) {
	err := data.Save()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to save %T: %w", data, err))
	}
	fmt.Printf("%T saved successfully!\n", data)
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
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	if value == "" {
		return "", errors.New("Input cannot be empty")
	}
	return strings.TrimSuffix(strings.TrimSuffix(value, "\n"), "\r"), nil
}
