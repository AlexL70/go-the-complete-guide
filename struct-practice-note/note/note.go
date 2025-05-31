package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Printf("\n\nMy Note:\nTitle: %s\nContent: %s\nCreated At: %s\n", n.title, n.content, n.createdAt.Format(time.RFC1123))
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	jsonContent, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonContent, 0644)
}
