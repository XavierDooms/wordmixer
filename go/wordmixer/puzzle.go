package wordmixer

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	x  int
	y  int
	xd int
	yd int
}

type Puzzle struct {
	word    *Word
	width   int
	height  int
	Board   [][]rune
	count   int
	iter    int
	sol     Solution
	success bool
}

func NewPuzzle(word string, size float32) *Puzzle {
	height := float32(len(word)) * size
	p := Puzzle{word: NewWord(word), width: 5 + int(height*1.5), height: 5 + int(height)}
	p.Board = generateRandomBoard(&p)
	for true {
		p.iter++
		/* if p.iter > 10 {
			return &p
		} */
		p.Board = generateRandomBoard(&p)
		countMatches(&p)
		if p.count == 0 {
			insertRandomMatche(&p)
			countMatches(&p)
		}
		if p.count == 1 {
			break
		}
	}
	fmt.Println("Iterations", p.iter)
	p.success = true
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

func countMatches(p *Puzzle) {
	p.count = 0
	for i, _ := range p.Board {
		checkBoardRow(p, i)
	}
	fmt.Println("Counted ", p.count)
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
				p.sol.x = i
				p.sol.y = j
				p.sol.xd = x
				p.sol.yd = y
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

func SolutionToString(p *Puzzle) string {
	solution := fmt.Sprint("x=", p.sol.y+1, " y=", p.sol.x+1)
	if p.sol.yd < 0 {
		solution += " left"
	} else if p.sol.yd > 0 {
		solution += " right"
	}
	if p.sol.xd < 0 {
		solution += " upper"
	} else if p.sol.xd > 0 {
		solution += " down"
	}
	return solution
}

func insertRandomMatche(p *Puzzle) {
	fmt.Println("h=", p.height, " w=", p.width)
	s := Solution{
		x:  rand.Intn(p.height - 1),
		y:  rand.Intn(p.width - 1),
		xd: 0,
		yd: 0,
	}
	fmt.Println("s ", s)

	if s.x-p.word.len < 0 {
		s.xd = rand.Intn(1)
	} else if s.x+p.word.len > p.height {
		s.xd = rand.Intn(1) - 1
	} else {
		s.xd = rand.Intn(2) - 1
	}

	if s.xd == 0 {
		fmt.Println("zero")
		if s.y-p.word.len < 0 {
			s.yd = 1
		} else if s.y+p.word.len > p.width {
			s.yd = -1
		} else {
			s.yd = rand.Intn(1)
			if s.yd == 0 {
				s.yd = -1
			}
		}
	} else if s.y-p.word.len < 0 {
		s.yd = rand.Intn(1)
	} else if s.y+p.word.len > p.width {
		s.yd = rand.Intn(1) - 1
	} else {
		s.yd = rand.Intn(2) - 1
	}

	fmt.Println("s ", s)
	for i, r := range p.word.runes {
		p.Board[s.x+i*s.xd][s.y+i*s.yd] = r
		fmt.Println("setting r=", string(r), " x=", s.x+i*s.xd, " y=", s.y+i*s.yd)
	}
	fmt.Println("s ", s)
}
