package main

import (
	"fmt"

	"../wordmixer"
)

func main() {
	word := wordmixer.NewWord("test")
	println(word.Word)

	puzzle := wordmixer.NewPuzzle("gdpr", 15, 10)
	fmt.Println(puzzle.Board)
}
