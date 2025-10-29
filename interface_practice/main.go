package main

import (
	"bufio"
	"fmt"
	"interface_practice/note"
	"interface_practice/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

//type displayer interface {
//	Display() error
//}

// embedded interface, like mixing in, kind of...
type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")
	todoThing, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}

	err = outputData(todoThing)

	if err != nil {
		return
	}

}

func getNoteData() (string, string) {

	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		return err
	}

	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// can use any as type, it's an alias for interface{}
func printSomething(value interface{}) {
	//to match on types:
	switch value.(type) {
	case string:
		fmt.Println(value)
	case int:
		fmt.Println("Integer", value)
	case float64:
		fmt.Println("Float", value)
		//default:
		//...
	}
	fmt.Println(value)
}

func printSomethingElse(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}
	stringVal, ok := value.(string)
	if ok {
		fmt.Println("Float:", stringVal)
		return
	}
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
