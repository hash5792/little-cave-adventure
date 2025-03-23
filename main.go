// https://helderman.github.io/htpataic/htpataic03.html

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hash5792/little-cave-adventure/parser"
)

var input string = "look around"

func getInput() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\n--> ")

	input, _ = reader.ReadString('\n')
}

func main() {
	fmt.Printf("Welcome to Little Cave Adventure.\n")

	for ; parser.ParseAndExecute(input); getInput() {
	}

	fmt.Printf("\nBye!\n")
}
