package greetings

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func Hello(name string) (message string, err error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message = fmt.Sprintf("Hi, %v. Welcome!", name)

	file, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return message, err
	}
	defer file.Close() // закрыть файл при выходе из функции

	log.SetOutput(file)
	log.Println("Log bla-blala")

	return message, nil
}
