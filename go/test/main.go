package main

import (
	"../wordmixer"
)

func main() {
	word := wordmixer.NewWord("test")
	println(word.Word)

	puzzle := wordmixer.NewPuzzle("gdpr", 15, 10)
	wordmixer.PrintBoard(puzzle)
}
