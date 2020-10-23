package wordmixer

import (
	"math/rand"
	"strings"
	"time"
)

type Word struct {
	Word  string
	runes []rune
	len   int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewWord(w string) *Word {
	wm := Word{Word: w}
	wm.runes = []rune(strings.ToUpper(w))
	wm.len = len(wm.runes)
	return &wm
}

func RandomWordRune(word *Word) rune {
	random_index := rand.Intn(word.len)
	return word.runes[random_index]
}
