package main

import (
	"../wordmixer"
)

func main() {
	word := wordmixer.NewWord("test")
	println(word.Word)

	puzzle := wordmixer.NewPuzzle("testvarken", 1.5)
	wordmixer.PrintBoard(puzzle)

	//println(wordmixer.ToHtml(puzzle))
}
