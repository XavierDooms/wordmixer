package wordmixer

type word struct {
	word string
}

func newWord(w string) *word {
	wm := word{w}
	return &wm
}
