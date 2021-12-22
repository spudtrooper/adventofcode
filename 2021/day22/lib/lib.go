package lib

import (
	"log"
	"regexp"
	"strings"

	"github.com/spudtrooper/adventofcode/common"
	"github.com/spudtrooper/adventofcode/common/must"
	"github.com/thomaso-mirodin/intmath/intgr"
)

var (
	// on x=-20..26,y=-36..17,z=-47..7
	cuboidRE = regexp.MustCompile(`(on|off) x=([\-\+]?\d+)\.\.([\-\+]?\d+),y=([\-\+]?\d+)\.\.([\-\+]?\d+),z=([\-\+]?\d+)\.\.([\-\+]?\d+)`)
)

type rng struct {
	from, to int
}

func (r rng) Min() int { return intgr.Min(r.from, r.to) }
func (r rng) Max() int { return intgr.Max(r.from, r.to) }
func (r rng) Len() int { return intgr.Abs(r.from - r.to) }

type cuboid struct {
	on      bool
	x, y, z rng
}

func Part1FromString(input string) int {
	var cuboids []cuboid
	for _, m := range cuboidRE.FindAllStringSubmatch(input, -1) {
		action, x1, x2, y1, y2, z1, z2 := m[1], must.Atoi(m[2]), must.Atoi(m[3]), must.Atoi(m[4]), must.Atoi(m[5]), must.Atoi(m[6]), must.Atoi(m[7])
		cuboids = append(cuboids, cuboid{action == "on", rng{x1, x2}, rng{y1, y2}, rng{z1, z2}})
	}

	for i, c := range cuboids {
		log.Printf("cuboids[%d]: %v", i, c)
	}

	board := common.MakeEmptyBoolBoard3D(101, 101, 101)
	for _, c := range cuboids {
		if (c.x.Min() < -50 || c.x.Max() > 50) ||
			(c.y.Min() < -50 || c.y.Max() > 50) ||
			(c.z.Min() < -50 || c.z.Max() > 50) {
			continue
		}
		for x := c.x.Min(); x <= c.x.Max(); x++ {
			for y := c.y.Min(); y <= c.y.Max(); y++ {
				for z := c.z.Min(); z <= c.z.Max(); z++ {
					board.Set(y+50, x+50, z+50, c.on)
				}
			}
		}
	}

	cnt := 0
	board.Traverse(func(x, y, z int, v bool) {
		if v {
			cnt++
		}
	})
	return cnt
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
