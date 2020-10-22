package main

import (
	"../wordmixer"
)

func main() {
	word := wordmixer.NewWord("test")
	println(word.Word)
}
