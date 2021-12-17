package lib

import (
	"log"
	"math"
	"regexp"

	"github.com/spudtrooper/adventofcode/common/geom"
	"github.com/spudtrooper/adventofcode/common/ints"
	"github.com/spudtrooper/adventofcode/common/must"
)

var (
	// target area: x=20..30, y=-10..-5
	inputRE = regexp.MustCompile(`target area: x=([\-\+]?\d+)\.\.([\-\+]?\d+), y=([\-\+]?\d+)\.\.([\-\+]?\d+)`)
)

type velocity geom.Point
type position geom.Point
type probe struct {
	vel velocity
	pos position
}

type target geom.Rect

func findHighestY(t target, vel velocity) (highest int, hitTarget bool) {
	p := probe{vel: vel, pos: geom.MakePoint(0, 0)}
	highest = math.MinInt
	for {
		if p.pos.Inside(t) {
			hitTarget = true
		}
		highest = ints.Max(highest, p.pos.Y())
		p.pos = p.pos.MoveBy(p.vel)
		var dx int
		if p.vel.X() > 0 {
			dx = -1
		} else if p.vel.X() < 0 {
			dx = 1
		}
		p.vel = p.vel.Move(dx, -1)
		if (p.vel.X() > 0 && p.pos.X() > t.SE().X()) ||
			(p.vel.X() < 0 && p.pos.X() < t.NW().X()) ||
			(p.vel.Y() < 0 && p.pos.Y() < t.SE().Y()) {
			break
		}
	}
	return
}

func Part1FromString(input string) int {
	m := inputRE.FindStringSubmatch(input)
	t := geom.MakeRect(must.Atoi(m[1]), must.Atoi(m[2]), must.Atoi(m[3]), must.Atoi(m[4]))

	highest := math.MinInt
	for vx := -t.NW().X(); vx <= t.SE().X(); vx++ {
		for vy := t.SE().Y(); vy <= -t.SE().Y(); vy++ {
			vel := geom.MakePoint(vx, vy)
			if height, hitTarget := findHighestY(t, vel); hitTarget {
				highest = ints.Max(highest, height)
			}
		}
	}

	return highest
}

func Part2FromString(input string) int {
	m := inputRE.FindStringSubmatch(input)
	t := geom.MakeRect(must.Atoi(m[1]), must.Atoi(m[2]), must.Atoi(m[3]), must.Atoi(m[4]))

	vels := 0
	for vx := -t.NW().X(); vx <= t.SE().X(); vx++ {
		for vy := t.SE().Y(); vy <= -t.SE().Y(); vy++ {
			vel := geom.MakePoint(vx, vy)
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
