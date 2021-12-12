package day09

import (
	"log"
	"sort"

	"github.com/spudtrooper/adventofcode/common/must"
)

type heightmap []row
type row []int

func (h heightmap) Dims() (width int, height int) {
	height, width = len(h), len(h[0])
	return
}

func (h heightmap) Adj(x, y, or int) (left, right, top, bottom int) {
	width, height := h.Dims()
	left, right, top, bottom = or, or, or, or
	if x > 0 {
		left = h[y][x-1]
	}
	if x < width-1 {
		right = h[y][x+1]
	}
	if y > 0 {
		top = h[y-1][x]
	}
	if y < height-1 {
		bottom = h[y+1][x]
	}
	return
}

func (h heightmap) LowPoints(f func(x, y int)) {
	isLowPoint := func(v, x, y int) bool {
		left, right, top, bottom := h.Adj(x, y, 10)
		return v < left && v < right && v < top && v < bottom
	}

	for y, row := range h {
		for x, v := range row {
			if isLowPoint(v, x, y) {
				f(x, y)
			}
		}
	}
}

func Part1(input string) int {
	var h heightmap
	for _, line := range must.ReadLines(input) {
		row := must.SplitInts(line, "")
		h = append(h, row)
	}

	var res int
	h.LowPoints(func(x, y int) {
		res += 1 + h[y][x]
	})

	return res
}

type pointSet map[int]bool

func pointHash(h heightmap, x, y int) int {
	width, _ := h.Dims()
	return y*width + x
}

func findBasin(h heightmap, x, y int, basin pointSet) {
	hash := pointHash(h, x, y)
	if _, searched := basin[hash]; searched {
		return
	}

	v := h[y][x]
	left, right, top, bottom := h.Adj(x, y, -1)

	if left > v {
		findBasin(h, x-1, y, basin)
	}
	if right > v {
		findBasin(h, x+1, y, basin)
	}
	if top > v {
		findBasin(h, x, y-1, basin)
	}
	if bottom > v {
		findBasin(h, x, y+1, basin)
	}

	if v != 9 {
		basin[hash] = true
	}
}

func Part2(input string) int {
	var h heightmap
	for _, line := range must.ReadLines(input) {
		row := must.SplitInts(line, "")
		h = append(h, row)
	}

	var basinSizes []int
	h.LowPoints(func(x, y int) {
		basin := pointSet{}
		findBasin(h, x, y, basin)
		basinSizes = append(basinSizes, len(basin))
	})

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	log.Printf("basinSizes: %v", basinSizes)
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}
