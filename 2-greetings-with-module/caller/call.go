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

	message, err := greetings.Hello("Go")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
