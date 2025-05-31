package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	jsonContent, err := json.MarshalIndent(todo, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal todo: %w", err)
	}
	return os.WriteFile(fileName, jsonContent, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("todo content cannot be empty")
	}
	return Todo{Text: content}, nil
}
