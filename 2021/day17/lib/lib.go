package lib

import (
	"log"
	"math"
	"regexp"

	"github.com/spudtrooper/adventofcode/common/ints"
	"github.com/spudtrooper/adventofcode/common/must"
)

var (
	// target area: x=20..30, y=-10..-5
	inputRE = regexp.MustCompile(`target area: x=([\-\+]?\d+)\.\.([\-\+]?\d+), y=([\-\+]?\d+)\.\.([\-\+]?\d+)`)
)

type coord struct {
	x, y int
}

type velocity struct {
	x, y int
}

type probe struct {
	vel velocity
	pos coord
}

type target struct {
	nw, se coord
}

func findHighestY(t target, vel velocity) (highest int, hitTarget bool) {
	p := probe{vel: vel}
	highest = math.MinInt
	for {
		if x, y := p.pos.x, p.pos.y; x >= t.nw.x && x <= t.se.x && y <= t.nw.y && y >= t.se.y {
			hitTarget = true
		}
		highest = ints.Max(highest, p.pos.y)
		p.pos.x += p.vel.x
		p.pos.y += p.vel.y
		if p.vel.x > 0 {
			p.vel.x--
		} else if p.vel.x < 0 {
			p.vel.x++
		}
		p.vel.y--
		if (p.vel.x > 0 && p.pos.x > t.se.x) ||
			(p.vel.x < 0 && p.pos.x < t.nw.x) ||
			(p.vel.y < 0 && p.pos.y < t.se.y) {
			break
		}
	}
	return
}

func Part1FromString(input string) int {
	m := inputRE.FindStringSubmatch(input)
	minX, maxX, minY, maxY := must.Atoi(m[1]), must.Atoi(m[2]), must.Atoi(m[3]), must.Atoi(m[4])
	t := target{nw: coord{x: minX, y: maxY}, se: coord{x: maxX, y: minY}}

	highest := math.MinInt
	for vx := 1; vx < t.se.x; vx++ {
		for vy := 1; vy < -minY; vy++ {
			vel := velocity{x: vx, y: vy}
			if height, hitTarget := findHighestY(t, vel); hitTarget {
				highest = ints.Max(highest, height)
			}
		}
	}

	return highest
}

func Part2FromString(input string) int {
	m := inputRE.FindStringSubmatch(input)
	minX, maxX, minY, maxY := must.Atoi(m[1]), must.Atoi(m[2]), must.Atoi(m[3]), must.Atoi(m[4])
	t := target{nw: coord{x: minX, y: maxY}, se: coord{x: maxX, y: minY}}

	vels := 0
	for vx := -t.nw.x; vx <= t.se.x; vx++ {
		for vy := minY; vy <= -minY; vy++ {
			vel := velocity{x: vx, y: vy}
			if _, hitTarget := findHighestY(t, vel); hitTarget {
				log.Printf("vel: %v", vel)
				vels++
			}
		}
	}

	return vels
}

func Part1(input string) int {
	return Part1FromString(must.ReadAllFile(input))
}

func Part2(input string) int {
	return Part2FromString(must.ReadAllFile(input))
}
