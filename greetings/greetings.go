package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("name was empty")
	}

	message := fmt.Sprintf("Hi, %v. Welcome yo!", name)
	//message := fmt.Sprint(randomFormat()) //-- to make it fail tests
	return message, nil
}

func RandomHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name was empty")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome yo!",
		"Great to see you %v!",
		"Nice to see you %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}
