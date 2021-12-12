package lib

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/spudtrooper/adventofcode/common/must"
)

type octopuses [][]int

func (o octopuses) String() string {
	var (
		out    []string
		white  = color.New(color.FgWhite).Add(color.Bold)
		yellow = color.New(color.FgYellow).Add(color.Bold)
	)
	for _, row := range o {
		var r []string
		for _, x := range row {
			if x == 10 {
				r = append(r, yellow.Sprintf("%02d", x))
			} else if x > 10 {
				r = append(r, white.Sprintf("%02d", x))
			} else {
				r = append(r, fmt.Sprintf("%02d", x))
			}
		}
		out = append(out, strings.Join(r, " "))
	}
	return strings.Join(out, "\n")
}

func (o octopuses) Dims() (width int, height int) {
	height, width = len(o), len(o[0])
	return
}

func (o octopuses) Step() (int, octopuses) {
	new := append(octopuses{}, o...)

	for y, row := range new {
		for x := range row {
			new[y][x]++
		}
	}

	deltas := []struct {
		dx, dy int
	}{
		{dx: -1, dy: +0},
		{dx: -1, dy: +1},
		{dx: +0, dy: +1},
		{dx: +1, dy: +1},
		{dx: +1, dy: +0},
		{dx: +1, dy: -1},
		{dx: +0, dy: -1},
		{dx: -1, dy: -1},
	}

	flashes := 0
	for {
		newFlashes := 0
		for y, row := range new {
			for x := range row {
				if new[y][x] < 10 {
					continue
				}
				new[y][x] = 0
				newFlashes++
				for _, d := range deltas {
					nx, ny := x+d.dx, y+d.dy
					if w, h := new.Dims(); nx < 0 || ny < 0 || nx >= w || ny >= h {
						continue
					}
					if new[ny][nx] == 0 {
						continue
					}
					new[ny][nx]++
				}
			}
		}
		flashes += newFlashes
		if newFlashes == 0 {
			break
		}
	}

	return flashes, new
}

func Part1(input string) int {
	var o octopuses
	for _, line := range must.ReadLines(input) {
		o = append(o, must.SplitInts(line, ""))
	}

	log.Printf("Initially:\n%v\n", o)

	var total int
	for i := 0; i < 100; i++ {
		flashes, new := o.Step()
		total += flashes
		o = new
	}

	log.Printf("Finally:\n%v\n", o)

	return total
}

func Part2(input string) int {
	var o octopuses
	for _, line := range must.ReadLines(input) {
		o = append(o, must.SplitInts(line, ""))
	}

	w, h := o.Dims()
	for step := 1; ; step++ {
		if f, _ := o.Step(); f == w*h {
			return step
		}
	}
}
