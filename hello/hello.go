package main

import (
	"fmt"
	"log"

	"github.com/daniil-novel/greetings"
)

func main() {
	message, err := greetings.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
