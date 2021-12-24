package common

import (
	"strings"
)

type StringBoard interface {
	Dims() (width int, height int)
	Get(y, x int) string
	Set(y, x int, s string)
	Traverse(func(y, x int, s string))
	VisitNeighbors(x, y int, f func(x, y int, s string))
	Clone() StringBoard
	String() string
}

func ParseStringBoard(input string) StringBoard {
	var b stringBoard
	for _, line := range strings.Split(input, "\n") {
		b = append(b, strings.Split(line, ""))
	}
	return b
}

func MakeStringBoard(width, height int, char func(y, x int) string) StringBoard {
	var b stringBoard
	for y := 0; y < height; y++ {
		var row []string
		for x := 0; x < width; x++ {
			c := char(y, x)
			row = append(row, c)
		}
		b = append(b, row)
	}
	return b
}

func StringBoardIdentFn(s string) func(y, x int) string {
	return func(y, x int) string {
		return s
	}
}

func (b stringBoard) Traverse(f func(y, x int, s string)) {
	for y, row := range b {
		for x, s := range row {
			f(y, x, s)
		}
	}
}

func (b stringBoard) VisitNeighbors(x, y int, f func(x, y int, s string)) {
	w, h := b.Dims()
	try := func(x, y int) {
		if x < 0 || x >= w || y < 0 || y >= h {
			return
		}
		f(x, y, b[y][x])
	}
	try(x, y-1)
	// try(x+1, y-1)
	try(x+1, y)
	// try(x+1, y+1)
	try(x, y+1)
	// try(x-1, y+1)
	try(x-1, y)
	// try(x-1, y-1)
}

type stringBoard [][]string

func (b stringBoard) Dims() (width int, height int) {
	height, width = len(b), len(b[0])
	return
}

func (b stringBoard) Get(y, x int) string {
	return b[y][x]
}

func (b stringBoard) Set(y, x int, s string) {
	b[y][x] = s
}

func (b stringBoard) String() string {
	var lines []string
	for _, row := range b {
		lines = append(lines, strings.Join(row, ""))
	}
	return strings.Join(lines, "\n")
}

func (b stringBoard) Clone() StringBoard {
	var res stringBoard
	for _, r := range b {
		var row []string
		row = append(row, r...)
		res = append(res, row)
	}
	return res
}
