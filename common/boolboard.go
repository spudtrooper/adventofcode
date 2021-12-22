package common

import (
	"strings"
)

type BoolBoard interface {
	Dims() (width int, height int)
	Get(y, x int) bool
	Set(y, x int, s bool)
	Traverse(func(y, x int, s bool))
}

func MakeBoolBoard(width, height int, ctor func(y, x int) bool) BoolBoard {
	var b boolBoard
	for y := 0; y < height; y++ {
		var row []bool
		for x := 0; x < width; x++ {
			c := ctor(y, x)
			row = append(row, c)
		}
		b = append(b, row)
	}
	return b
}

func BoolBoardIdentFn(v bool) func(y, x int) bool {
	return func(y, x int) bool {
		return v
	}
}

func (b boolBoard) Traverse(f func(y, x int, s bool)) {
	for y, row := range b {
		for x, v := range row {
			f(y, x, v)
		}
	}
}

type boolBoard [][]bool

func (b boolBoard) Dims() (width int, height int) {
	height, width = len(b), len(b[0])
	return
}

func (b boolBoard) Get(y, x int) bool {
	return b[y][x]
}

func (b boolBoard) Set(y, x int, v bool) {
	b[y][x] = v
}

func (b boolBoard) String() string {
	var lines []string
	for _, row := range b {
		var bs []string
		for _, b := range row {
			v := "F"
			if b {
				v = "T"
			}
			bs = append(bs, v)
		}
		lines = append(lines, strings.Join(bs, ""))
	}
	return strings.Join(lines, "\n")
}
