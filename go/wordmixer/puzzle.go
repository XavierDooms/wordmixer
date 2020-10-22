package wordmixer

type Puzzle struct {
	word   *Word
	width  int
	height int
	Board  [][]rune
}

func NewPuzzle(word string, width int, height int) *Puzzle {
	p := Puzzle{word: NewWord(word), width: width, height: height}
	p.Board = generateRandomBoard(&p)
	for true {
		p.Board = generateRandomBoard(&p)
		if isValidBoard(p.Board) {
			break
		}
	}
	return &p
}

func generateRandomBoard(p *Puzzle) [][]rune {
	var b = make([][]rune, p.height)
	for i := 0; i < p.height; i++ {
		b[i] = make([]rune, p.width)
		for j := 0; j < p.width; j++ {
			b[i][j] = RandomWordRune(p.word)
		}
	}
	return b
}

func isValidBoard(b [][]rune) bool {
	return true
}
