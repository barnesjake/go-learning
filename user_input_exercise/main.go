package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"user_input_exercise/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed", err)
	}

	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {

	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content

}

func getUserInput(promptText string) string {
	fmt.Printf("%v ", promptText)

	//Stdin is just command line
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // on windows it can be represented as \r

	return text
}
