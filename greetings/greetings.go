package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (message string, err error) {
	// Return a greeting that embeds the name in a message.

	if message == "" {
		err = errors.New("empty message")
		err1 := fmt.Errorf("User message: %s is empty", message)
		return message, err1
	}

	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, err
}
