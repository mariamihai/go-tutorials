package greetings

import (
	"regexp"
	"testing"
)

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")

	if message != "" || err == nil {
		t.Fatalf(`Hello("") = [%q, %v], wants match for ["", error]`, message, err)
	}
}

func TestHelloName(t *testing.T) {
	name := "Go"
	want := regexp.MustCompile(`\b` + name + `\b`)

	message, err := Hello(name)

	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello(%q) = [%q, %v], wants match for [%#q, nil]`, name, message, err, want)
	}
}

func TestHellosEmpty(t *testing.T) {
	names := []string{"Go", ""}

	messages, err := Hellos(names)

	if messages != nil || err == nil {
		t.Fatalf(`Hello(%q) = [%q, %v], wants match for ["", error]`, names, messages, err)
	}
}

func TestHellosNames(t *testing.T) {
	names := []string{"Go", "Other Name"}
	messages, err := Hellos(names)

	if messages == nil || err != nil {
		t.Fatalf(`Hello(%q) = [%q, %v], wants match for ["", error]`, names, messages, err)
	}

	for _, name := range names {
		want := regexp.MustCompile(`\b` + name + `\b`)

		if !want.MatchString(messages[name]) || err != nil {
			t.Fatalf(`Hello(%q) = [%q, %v], wants match for [%#q, nil]`, name, messages[name], err, want)
		}
	}
}
