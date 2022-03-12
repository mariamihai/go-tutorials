package main

import (
	"fmt"
	"log"
	"random-prefix/greetings"
)

func main() {
	log.SetPrefix("greetings: ")

	// No timestamp, source file, line number
	log.SetFlags(0)

	// Get one greeting
	message, err := greetings.Hello("Go")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	// Get multiple greetings
	messages, err := greetings.Hellos([]string{"Go", "Other Name", "Another Name"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
