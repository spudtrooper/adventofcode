package day11

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/spudtrooper/adventofcode/common/must"
)

type octopuses []row
type row []int

func (o octopuses) String() string {
	var buf bytes.Buffer
	white := color.New(color.FgWhite).Add(color.Bold)
	for _, row := range o {
		for _, x := range row {
			if x == 0 {
				white.Fprintf(&buf, "%d", x)
			} else {
				fmt.Fprintf(&buf, "%d", x)
			}
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func (o octopuses) Debug() string {
	var out []string
	white := color.New(color.FgWhite).Add(color.Bold)
	yellow := color.New(color.FgYellow).Add(color.Bold)
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
	var new octopuses
	new = append(new, o...)

	for y, row := range new {
		for x := range row {
			new[y][x]++
		}
	}

	log.Printf("After initial step:\n%v\n", new.Debug())

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
		startFlashes := flashes
		for y, row := range new {
			for x := range row {
				if new[y][x] < 10 {
					continue
				}
				new[y][x] = 0
				flashes++
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
		if startFlashes == flashes {
			break
		}
	}
	log.Printf("After check(s):\n%v\n", new.Debug())

	return flashes, new
}

func Part1(input string) int {
	var o octopuses
	for _, line := range must.ReadStrings(input) {
		row := must.SplitInts(line, "")
		o = append(o, row)
	}

	log.Printf("Initially:\n%s\n", o.Debug())

	var total int
	for i := 0; i < 100; i++ {
		flashes, new := o.Step()
		total += flashes
		o = new
	}

	log.Printf("Finally:\n%s\n", o.Debug())

	return total
}

func Part2(input string) int {
	// TODO
	return -1
}
