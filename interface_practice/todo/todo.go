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

	//json.Marshal returns [] already
	// json.Marshal needs struct data that is publicly available/not exclusive to the package
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	//WriteFile can return error so return result of that
	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("Invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}
