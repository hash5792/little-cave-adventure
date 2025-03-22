package parser

import (
	"fmt"
	"strings"
)

func ParseAndExecute(input string) bool {

	// Basic parsing by splitting the input with space
	// We assume first word is verb and second word is noun
	var words []string = strings.Split(input, " ")

	var verb string = words[0]
	// var noun string = words[1] // Not used for the moment

	switch verb {
	case "quit":
		return false
	case "look":
		fmt.Printf("It is very dark in here.\n")
	case "go":
		fmt.Printf("It's too dark to go anywhere.\n")
	default:
		fmt.Printf("I don't know how to %v.\n", verb)
	}

	return true
}
