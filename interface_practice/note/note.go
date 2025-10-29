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
	Title     string    `json:"title"` //struct tags can be used to set metadata, i.e. we can control json key format etc. we could have `json:"abc"` to make the key value abc
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	//json.Marshal returns [] already
	// json.Marshal needs struct data that is publicly available/not exclusive to the package
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	//WriteFile can return error so return result of that
	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
