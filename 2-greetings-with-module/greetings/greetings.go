package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomGreeting(), name)

	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomGreeting() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Welcome, %v!",
	}

	return formats[rand.Intn(len(formats))]
}
