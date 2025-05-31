package note

import (
	"errors"
	"fmt"
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
