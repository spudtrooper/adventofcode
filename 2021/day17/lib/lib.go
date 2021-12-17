package lib

import (
	"log"
	"math"
	"regexp"

	"github.com/spudtrooper/adventofcode/common"
	"github.com/spudtrooper/adventofcode/common/ints"
	"github.com/spudtrooper/adventofcode/common/must"
)

var (
	// target area: x=20..30, y=-10..-5
	inputRE = regexp.MustCompile(`target area: x=([\-\+]?\d+)\.\.([\-\+]?\d+), y=([\-\+]?\d+)\.\.([\-\+]?\d+)`)
)

type velocity struct {
	x, y int
}

type probe struct {
	vel velocity
	pos common.Point
}

type target common.Rect

func findHighestY(t target, vel velocity) (highest int, hitTarget bool) {
	p := probe{vel: vel, pos: common.MakePoint(0, 0)}
	highest = math.MinInt
	for {
		if p.pos.Inside(t) {
			hitTarget = true
		}
		highest = ints.Max(highest, p.pos.Y())
		p.pos = p.pos.Move(p.vel.x, p.vel.y)
		if p.vel.x > 0 {
			p.vel.x--
		} else if p.vel.x < 0 {
			p.vel.x++
		}
		p.vel.y--
		if (p.vel.x > 0 && p.pos.X() > t.SE().X()) ||
			(p.vel.x < 0 && p.pos.X() < t.NW().X()) ||
			(p.vel.y < 0 && p.pos.Y() < t.SE().Y()) {
			break
		}
	}
	return
}

func Part1FromString(input string) int {
	m := inputRE.FindStringSubmatch(input)
	t := common.MakeRect(must.Atoi(m[1]), must.Atoi(m[2]), must.Atoi(m[3]), must.Atoi(m[4]))

	highest := math.MinInt
	for vx := -t.NW().X(); vx <= t.SE().X(); vx++ {
		for vy := t.SE().Y(); vy <= -t.SE().Y(); vy++ {
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
	t := common.MakeRect(must.Atoi(m[1]), must.Atoi(m[2]), must.Atoi(m[3]), must.Atoi(m[4]))

	vels := 0
	for vx := -t.NW().X(); vx <= t.SE().X(); vx++ {
		for vy := t.SE().Y(); vy <= -t.SE().Y(); vy++ {
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
