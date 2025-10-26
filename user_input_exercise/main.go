package main

import (
	"errors"
	"fmt"
	"user_input_exercise/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData() (string, string) {

	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content

}

func getUserInput(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)

	return value
}
