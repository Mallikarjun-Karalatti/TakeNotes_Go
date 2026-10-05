package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Time    time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content cannot be empty")
	}
	return &Note{
		Title:   title,
		Content: content,
		Time:    time.Now(),
	}, nil
}

func (n *Note) SaveToFile() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ReplaceAll(fileName, "/", "_")
	fileName = strings.ToLower(fileName)
	fileName = fileName + ".json"

	content, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, content, 0644)
}
