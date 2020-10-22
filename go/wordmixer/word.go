package wordmixer

type Word struct {
	Word string
}

func NewWord(w string) *Word {
	wm := Word{w}
	return &wm
}
