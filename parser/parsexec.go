package parser

import (
	"fmt"
	"strings"

	"github.com/hash5792/little-cave-adventure/world"
)

func ParseAndExecute(input string) bool {

	// Basic parsing by splitting the input with space
	// We assume first word is verb and second word is noun
	var words []string = strings.Fields(input)

	var noun string
	var verb string = words[0]

	if len(words) > 1 {
		noun = words[1]
	}

	switch verb {
	case "quit":
		return false
	case "look":
		world.ExecuteLook(noun)
	case "go":
		world.ExecuteGo(noun)
	default:
		fmt.Printf("I don't know how to %v.\n", verb)
	}

	return true
}
