package lib

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/spudtrooper/adventofcode/common/must"
)

var (
	foldRE = regexp.MustCompile(`fold along ([xy])=(\d+)`)
)

type coord struct {
	x, y int
}
type fold struct {
	along string
	val   int
}
type paper struct {
	coords        map[int]coord
	width, height int
}

func (p paper) Debug(f fold) string {
	arr := make([][]string, p.height)
	for y := 0; y < p.height; y++ {
		arr[y] = make([]string, p.width)
		for x := range arr[y] {
			arr[y][x] = "."
			if f.along == "y" {
				if y == f.val {
					arr[y][x] = "-"
				}
			} else {
				if x == f.val {
					arr[y][x] = "|"
				}

			}
		}
	}
	var i int
	for _, c := range p.coords {
		if c.y < 0 || c.y >= p.height || c.x < 0 || c.x >= p.width {
			log.Fatalf("Invalid width=%d height=%d coords[%d]: %+v", p.width, p.height, i, c)
		}
		arr[c.y][c.x] = "#"
		i++
	}
	var buf bytes.Buffer
	for i, row := range arr {
		buf.WriteString(fmt.Sprintf("%3d %s", i, strings.Join(row, "")))
		buf.WriteString("\n")
	}
	return buf.String()
}

func Part1(input string) int {
	var folds []fold
	var coords []coord
	var maxX, maxY int

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	hash := func(p paper, c coord) int {
		return p.width*c.y + c.x
	}

	for _, line := range must.ReadLines(input) {
		if line == "" {
			continue
		}
		if m := foldRE.FindStringSubmatch(line); len(m) == 3 {
			along, val := m[1], must.Atoi(m[2])
			folds = append(folds, fold{along, val})
		} else {
			parts := must.SplitInts(line, ",")
			x, y := parts[0], parts[1]
			coords = append(coords, coord{x: x, y: y})
			maxX = max(maxX, x)
			maxY = max(maxY, y)
		}
	}

	initHeight, initWidth := maxY+1, maxX+1

	log.Printf("initWidth: %v", initWidth)
	log.Printf("initHeight: %v", initHeight)
	log.Printf("folds: %+v", folds)
	log.Printf("coords: %+v", coords)

	makePaper := func(height, width int) paper {
		return paper{height: height, width: width, coords: map[int]coord{}}
	}

	p := makePaper(initHeight, initWidth)
	for _, c := range coords {
		p.coords[hash(p, c)] = c
	}

	// log.Printf("paper\n%v", p.Debug())

	fold := func(p paper, f fold) paper {
		if f.along == "y" {
			newHeight, del := p.height/2, 1
			if p.height%2 == 0 {
				newHeight++
				del = 0
			}
			new := makePaper(newHeight, p.width)
			for _, c := range p.coords {
				if c.y > f.val {
					c.y = p.height - c.y - del
				}
				new.coords[hash(p, c)] = c
			}
			return new
		}

		newWidth, del := p.width/2, 1
		if p.width%2 == 0 {
			newWidth++
			del = 0
		}
		new := makePaper(p.height, newWidth)
		for _, c := range p.coords {
			if c.x > f.val {
				c.x = p.width - c.x - del
			}
			new.coords[hash(p, c)] = c
		}

		// log.Printf("paper\n%s", new.Debug())

		return new
	}

	p = fold(p, folds[0])

	return len(p.coords)
}

func Part2(input string) int {
	var folds []fold
	var coords []coord
	var maxX, maxY int

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	hash := func(p paper, c coord) int {
		return p.width*c.y + c.x
	}

	for _, line := range must.ReadLines(input) {
		if line == "" {
			continue
		}
		if m := foldRE.FindStringSubmatch(line); len(m) == 3 {
			along, val := m[1], must.Atoi(m[2])
			folds = append(folds, fold{along, val})
		} else {
			parts := must.SplitInts(line, ",")
			x, y := parts[0], parts[1]
			coords = append(coords, coord{x: x, y: y})
			maxX = max(maxX, x)
			maxY = max(maxY, y)
		}
	}

	initHeight, initWidth := maxY+1, maxX+1

	log.Printf("initWidth: %v", initWidth)
	log.Printf("initHeight: %v", initHeight)
	log.Printf("folds: %+v", folds)
	log.Printf("coords: %+v", coords)

	makePaper := func(height, width int) paper {
		return paper{height: height, width: width, coords: map[int]coord{}}
	}

	p := makePaper(initHeight, initWidth)
	for _, c := range coords {
		p.coords[hash(p, c)] = c
	}

	fold := func(p paper, f fold) paper {
		if f.along == "y" {
			new := makePaper(f.val, p.width)
			for _, c := range p.coords {
				if c.y > f.val {
					c.y = f.val - (c.y - f.val)
				}
				new.coords[hash(p, c)] = c
			}
			return new
		}

		new := makePaper(p.height, f.val)
		for _, c := range p.coords {
			if c.x > f.val {
				c.x = f.val - (c.x - f.val)
			}
			new.coords[hash(p, c)] = c
		}

		return new
	}

	for _, f := range folds {
		p = fold(p, f)
		log.Printf("\n%v", p.Debug(f))
	}

	return len(p.coords)
}
