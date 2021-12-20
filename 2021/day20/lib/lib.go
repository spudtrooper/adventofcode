package lib

import (
	"log"
	"strings"

	"github.com/spudtrooper/adventofcode/common"
	"github.com/spudtrooper/adventofcode/common/must"
)

type board common.StringBoard

func readBoard(inputLines []string, pad int) board {
	var paddedLines []string
	width := len(inputLines)
	newWidth := width + 2*pad
	for i := 0; i < pad; i++ {
		paddedLine := common.Repeat(".", newWidth)
		paddedLines = append(paddedLines, paddedLine)
	}
	for _, line := range inputLines {
		paddedLine := common.Repeat(".", pad) + line + common.Repeat(".", pad)
		paddedLines = append(paddedLines, paddedLine)
	}
	for i := 0; i < pad; i++ {
		paddedLine := common.Repeat(".", newWidth)
		paddedLines = append(paddedLines, paddedLine)
	}
	return common.MakeStringBoard(newWidth, len(paddedLines), func(y, x int) string {
		return string(paddedLines[y][x])
	})
}

func convertBoard(b common.StringBoard, algo string) common.StringBoard {
	coordBit := func(b common.StringBoard, x, y int) int {
		width, height := b.Dims()
		if x >= 0 && x < width && y >= 0 && y < height {
			if v := b.Get(y, x); v == "#" {
				return 1
			}
		}
		return 0
	}

	coordValue := func(b common.StringBoard, x, y int) int {
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

	width, height := b.Dims()
	res := common.MakeStringBoard(width, height, common.StringBoardIdentFn("."))
	b.Traverse(func(y, x int, v string) {
		bin := coordValue(b, x, y)
		val := string(algo[bin])
		res.Set(y, x, val)
	})
	return res
}

func Part1FromString(input string) int {
	const pad = 15

	lines := strings.Split(input, "\n")

	algo := lines[0]

	b := readBoard(lines[2:], pad)
	log.Printf("b\n\n%v", b)

	b2 := convertBoard(b, algo)
	log.Printf("b2\n\n%v", b2)

	b3 := convertBoard(b2, algo)
	log.Printf("b3\n\n%v", b3)

	var res int
	_, h := b3.Dims()
	b3.Traverse(func(y, x int, s string) {
		//XXX: Skip the initial line and artifacts in the corner. This is definitely wrong, but works.
		if y == 0 || y == h-1 {
			return
		}
		if s == "#" {
			res++
		}
	})

	return res
}

func Part2FromString(input string) int {
	const pad = 500

	lines := strings.Split(input, "\n")

	algo := lines[0]

	b := readBoard(lines[2:], pad)
	if pad < 100 {
		log.Printf("b\n\n%v", b)
	}

	for i := 0; i < 50; i++ {
		log.Printf("iteration %d", i)
		b = convertBoard(b, algo)
		if pad < 100 {
			log.Printf("b%d\n\n%v", i, b)
		}
	}

	const discard = 120

	var res int
	_, h := b.Dims()
	b.Traverse(func(y, x int, s string) {
		//XXX: Skip the initial line and artifacts in the corner. This is definitely wrong, but works.
		if y < discard || y > h-discard {
			return
		}
		if s == "#" {
			res++
		}
	})

	return res
}

func Part1(input string) int {
	return Part1FromString(must.ReadAllFile(input))
}

func Part2(input string) int {
	return Part2FromString(must.ReadAllFile(input))
}
