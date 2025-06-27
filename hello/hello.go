package main

import (
	"example/greetings"
	"fmt"
	"log"
)

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())

	simpleHello, _ := greetings.Hello("Gladys")
	fmt.Println(simpleHello)

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("123")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	randomMessage, err := greetings.RandomHello("Gladys")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(randomMessage)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Helena", "Amy", "Ems"}
	groupMessages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(groupMessages)

}
