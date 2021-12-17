package day05

import (
	"log"
	"regexp"

	"github.com/spudtrooper/adventofcode/common/ints"
	"github.com/spudtrooper/adventofcode/common/must"
	"github.com/thomaso-mirodin/intmath/intgr"
)

var (
	configLineRE = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
)

type line struct {
	x1, y1, x2, y2 int
}

func readConfig(input string) []line {
	var res []line
	for _, m := range configLineRE.FindAllStringSubmatch(must.ReadAllFile(input), -1) {
		x1, y1, x2, y2 := must.Atoi(m[1]), must.Atoi(m[2]), must.Atoi(m[3]), must.Atoi(m[4])
		res = append(res, line{x1: x1, y1: y1, x2: x2, y2: y2})
	}
	return res
}

func Part1(input string) int {
	f := ints.MakeBoard(1000, 1000)
	for _, line := range readConfig(input) {
		if line.x1 != line.x2 && line.y1 != line.y2 {
			continue
		}
		for y := intgr.Min(line.y1, line.y2); y <= intgr.Max(line.y1, line.y2); y++ {
			for x := intgr.Min(line.x1, line.x2); x <= intgr.Max(line.x1, line.x2); x++ {
				log.Printf("height: %d", len(f))
				log.Printf("width: %d", len(f[0]))
				f[y][x]++
			}
		}
	}

	var res int
	for _, row := range f {
		for _, v := range row {
			if v >= 2 {
				res++
			}
		}
	}
	return res
}

func Part2(input string) int {
	delta := func(a, b int) int {
		if a == b {
			return 0
		}
		if a > b {
			return -1
		}
		return 1
	}

	f := ints.MakeBoard(1000, 1000)
	for _, line := range readConfig(input) {
		for x, y := line.x1, line.y1; ; {
			f[y][x]++
			if x == line.x2 && y == line.y2 {
				break
			}
			x += delta(line.x1, line.x2)
			y += delta(line.y1, line.y2)
		}
	}

	var res int
	for _, row := range f {
		for _, v := range row {
			if v >= 2 {
				res++
			}
		}
	}
	return res
}
