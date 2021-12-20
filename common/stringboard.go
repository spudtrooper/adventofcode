package common

import "strings"

type StringBoard interface {
	Dims() (width int, height int)
	Get(y, x int) string
	Set(y, x int, s string)
	Traverse(func(y, x int, s string))
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
func (b stringBoard) Traverse(f func(y, x int, s string)) {
	for y, row := range b {
		for x, s := range row {
			f(y, x, s)
		}
	}
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
