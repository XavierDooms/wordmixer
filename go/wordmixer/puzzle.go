package wordmixer

import "fmt"

type Puzzle struct {
	word   *Word
	width  int
	height int
	Board  [][]rune
	count  int
	iter   int
}

func NewPuzzle(word string, width int, height int) *Puzzle {
	p := Puzzle{word: NewWord(word), width: width, height: height}
	p.Board = generateRandomBoard(&p)
	for true {
		p.iter++
		p.Board = generateRandomBoard(&p)
		if isValidBoard(&p) {
			break
		}
	}
	fmt.Println("Iterations", p.iter)
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

func isValidBoard(p *Puzzle) bool {
	p.count = 0
	for i, _ := range p.Board {
		checkBoardRow(p, i)
	}
	fmt.Println("Counted ", p.count)
	return p.count == 1
}

func checkBoardRow(p *Puzzle, i int) bool {
	for j, r := range p.Board[i] {
		if r != p.word.runes[0] {
			continue
		}
		checkBoardRune(p, i, j)
	}
	return true
}

func checkBoardRune(p *Puzzle, i int, j int) bool {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if (x == 0 && y == 0) ||
				(x < 0 && i-p.word.len < 0) ||
				(y < 0 && j-p.word.len < 0) ||
				(x > 0 && i+p.word.len > (p.height-1)) ||
				(y > 0 && j+p.word.len > (p.width-1)) {
				continue
			}
			match := true
			for k, r := range p.word.runes {
				if p.Board[i+k*x][j+k*y] == r {
					continue
				}
				match = false
				break
			}
			if match {
				p.count++
			}
		}
	}
	return true
}

func PrintBoard(p *Puzzle, params ...string) {
	fmt.Println(BoardToString(p, "\n"))
}

func BoardToString(p *Puzzle, end string) string {
	str := ""
	for _, row := range p.Board {
		for _, rune := range row {
			str += string(rune) + " "
		}
		str += end
	}
	return str
}
