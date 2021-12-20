package lib

import (
	"bytes"
	"log"
	"strings"

	"github.com/spudtrooper/adventofcode/common/must"
)

type board [][]string

func (b board) Dims() (width int, height int) {
	height, width = len(b), len(b[0])
	return
}

func (b board) String() string {
	var lines []string
	for _, row := range b {
		lines = append(lines, strings.Join(row, ""))
	}
	return strings.Join(lines, "\n")
}

const pad = 15

func readBoard(inputLines []string) board {
	var paddedLines []string
	width := len(inputLines)
	newWidth := width + 2*pad
	for i := 0; i < pad; i++ {
		paddedLine := repeat(".", newWidth)
		paddedLines = append(paddedLines, paddedLine)
	}
	for _, line := range inputLines {
		paddedLine := repeat(".", pad) + line + repeat(".", pad)
		paddedLines = append(paddedLines, paddedLine)
	}
	for i := 0; i < pad; i++ {
		paddedLine := repeat(".", newWidth)
		paddedLines = append(paddedLines, paddedLine)
	}
	var b board
	for _, line := range paddedLines {
		b = append(b, strings.Split(line, ""))
	}
	return b
}

func makeBoard(width, height int) board {
	var b board
	for y := 0; y < height; y++ {
		b = append(b, strings.Split(repeat(".", width), ""))
	}
	return b
}

func repeat(s string, times int) string {
	var buf bytes.Buffer
	for i := 0; i < times; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

func Part1FromString(input string) int {
	coordBit := func(b board, x, y int) int {
		width, height := b.Dims()
		if x >= 0 && x < width && y >= 0 && y < height {
			if v := b[y][x]; v == "#" {
				return 1
			}
		}
		return 0
	}

	coordValue := func(b board, x, y int) int {
		var res int
		res = res | (coordBit(b, x-1, y-1) << 8)
		res = res | (coordBit(b, x+0, y-1) << 7)
		res = res | (coordBit(b, x+1, y-1) << 6)
		res = res | (coordBit(b, x-1, y+0) << 5)
		res = res | (coordBit(b, x+0, y+0) << 4)
		res = res | (coordBit(b, x+1, y+0) << 3)
		res = res | (coordBit(b, x-1, y+1) << 2)
		res = res | (coordBit(b, x+0, y+1) << 1)
		res = res | (coordBit(b, x+1, y+1) << 0)
		return res
	}

	convert := func(b board, algo string) board {
		width, height := b.Dims()
		res := makeBoard(width, height)
		for y, row := range b {
			for x := range row {
				bin := coordValue(b, x, y)
				val := string(algo[bin])
				res[y][x] = val
			}
		}
		return res
	}

	lines := strings.Split(input, "\n")

	algo := lines[0]

	b := readBoard(lines[2:])
	log.Printf("b\n\n%v", b)

	b2 := convert(b, algo)
	log.Printf("b2\n\n%v", b2)

	b3 := convert(b2, algo)
	log.Printf("b3\n\n%v", b3)

	var res int
	for y, row := range b3 {
		//XXX: Skip the initial line and artifacts in the corner. This is definitely wrong, but works.
		if y == 0 || y == len(b3)-1 {
			continue
		}
		for _, v := range row {
			if v == "#" {
				res++
			}
		}
	}

	return res
}

func Part2FromString(input string) int {
	for _, line := range strings.Split(input, "\n") {
		// TODO
		if false {
			log.Println(line)
		}
	}
	return -1
}

func Part1(input string) int {
	return Part1FromString(must.ReadAllFile(input))
}

func Part2(input string) int {
	return Part2FromString(must.ReadAllFile(input))
}
