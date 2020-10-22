package main

import (
	"fmt"

	"../wordmixer"
)

func main() {
	word := wordmixer.NewWord("test")
	println(word.Word)

	puzzle := wordmixer.NewPuzzle("gdpr", 10, 15)
	fmt.Println(puzzle.Board)
}
