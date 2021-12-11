package day04

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/spudtrooper/adventofcode/common/must"
)

const (
	numRows = 5
	numCols = 5
)

type board struct {
	rows []*row
}

func (b *board) String() string {
	var rows []string
	for _, r := range b.rows {
		rows = append(rows, r.String())
	}
	return strings.Join(rows, "\n")
}

func (b *board) Mark(v int) {
	for _, r := range b.rows {
		for _, c := range r.cells {
			if c.val == v {
				c.marked = true
			}
		}
	}
}
func (b *board) CheckWinner() bool {
	for _, r := range b.rows {
		var sel int
		for _, c := range r.cells {
			if c.marked {
				sel++
			}
		}
		if sel == numCols {
			for _, c := range r.cells {
				c.winner = true
			}
			return true
		}
	}
	for col := 0; col < numCols; col++ {
		var selected []*cell
		for row := 0; row < numRows; row++ {
			if c := b.rows[row].cells[col]; c.marked {
				selected = append(selected, c)
			}
		}
		if len(selected) == numRows {
			for _, c := range selected {
				c.winner = true
			}
			return true
		}
	}
	return false
}

func (b *board) UnmarkedNumbers() int {
	var res int
	for _, r := range b.rows {
		for _, c := range r.cells {
			if !c.marked {
				res += c.val
			}
		}
	}
	return res
}

type row struct {
	cells []*cell
}

func (r *row) String() string {
	var cells []string
	for _, c := range r.cells {
		s := fmt.Sprintf("%2d", c.val)
		if c.winner {
			s = color.GreenString("%2d", c.val)
		} else if c.marked {
			s = color.YellowString("%2d", c.val)
		}
		cells = append(cells, s)
	}
	return strings.Join(cells, " ")
}

type cell struct {
	val    int
	marked bool
	winner bool
}

type game struct {
	numbers []int
	boards  []*board
}

func (g *game) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("All numbers: %v", g.numbers))
	for _, b := range g.boards {
		lines = append(lines, "\n", b.String())
	}
	return strings.Join(lines, "\n")
}

func readGame(input string) game {
	lines := must.ReadStrings(input)

	readNums := func(line, sep string) []int {
		var nums []int
		for _, s := range strings.Split(line, sep) {
			if s == sep || s == "" {
				continue
			}
			nums = append(nums, must.Atoi(s))
		}
		return nums
	}

	numbers := readNums(lines[0], ",")
	g := game{numbers: numbers}

	var b *board
	for _, line := range lines[1:] {
		if line == "" {
			b = &board{}
			g.boards = append(g.boards, b)
			continue
		}
		row := &row{}
		for _, n := range readNums(line, " ") {
			row.cells = append(row.cells, &cell{val: n})
		}
		b.rows = append(b.rows, row)
	}

	return g
}

func (g *game) FirstWinner() (*board, int) {
	for _, n := range g.numbers {
		for _, b := range g.boards {
			b.Mark(n)
			if b.CheckWinner() {
				return b, n
			}
		}
	}
	return nil, 0
}

func (g *game) LastWinner() (*board, int) {
	var lastWinner *board
	lastN := -1
	for _, n := range g.numbers {
		for _, b := range g.boards {
			if !b.CheckWinner() {
				b.Mark(n)
				if b.CheckWinner() {
					lastWinner = b
					lastN = n
				}
			}
		}
	}
	return lastWinner, lastN
}

func Part1(input string) int {
	g := readGame(input)

	if winner, n := g.FirstWinner(); winner != nil {
		log.Println()
		log.Println(g.String())
		log.Printf("winner\n%v", winner)
		unmarkedNumbers := winner.UnmarkedNumbers()
		mul := n * winner.UnmarkedNumbers()
		log.Printf("n: %d", n)
		log.Printf("unmarkedNumbers: %d", unmarkedNumbers)
		log.Printf("mul: %d", mul)
		return mul
	}

	return 0
}

func Part2(input string) int {
	g := readGame(input)

	if winner, n := g.LastWinner(); winner != nil {
		log.Println()
		log.Println(g.String())
		log.Printf("winner\n%v", winner)
		unmarkedNumbers := winner.UnmarkedNumbers()
		mul := n * unmarkedNumbers
		log.Printf("n: %d", n)
		log.Printf("unmarkedNumbers: %d", unmarkedNumbers)
		log.Printf("mul: %d", mul)
		return mul
	}

	return 0
}
