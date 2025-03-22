// https://helderman.github.io/htpataic/htpataic02.html

package main

import (
	"fmt"

	"github.com/hash5792/little-cave-adventure/parser"
)

var input string = "look around"

func getInput() {
	fmt.Printf("\n--> ")
	fmt.Scanln(&input)
}

func main() {
	fmt.Printf("Welcome to Little Cave Adventure.\n")

	for ; parser.ParseAndExecute(input); getInput() {
	}

	fmt.Printf("\nBye!\n")
}
