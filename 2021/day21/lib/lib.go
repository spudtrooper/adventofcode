package lib

import (
	"log"
	"regexp"
	"strings"

	"github.com/spudtrooper/adventofcode/common/must"
)

var (
	// Player 1 starting position: 4
	playerRE = regexp.MustCompile(`Player \d+ starting position: (\d+)`)
)

type player struct {
	id    int
	pos   int
	score int
}

type die struct {
	val int
}

func (d *die) next() int {
	res := d.val
	d.val = (d.val + 1) % 100
	return res
}

func (d *die) roll(p *player) {
	r1, r2, r3 := d.next()%100, d.next()%100, d.next()%100
	for i, n := 0, r1+r2+r3; i < n; i++ {
		p.pos++
		if p.pos == 11 {
			p.pos = 1
		}
	}
	p.score += p.pos
	log.Printf("Player %d rolls %d+%d+%d and moves to space %d for a total score of %d", p.id, r1, r2, r3, p.pos, p.score)
}

func Part1FromString(input string) int {
	lines := strings.Split(input, "\n")
	var p1, p2 player
	if m := playerRE.FindStringSubmatch(lines[0]); len(m) == 2 {
		pos := must.Atoi(m[1])
		p1 = player{id: 1, pos: pos}
	} else {
		log.Fatalf("invalid player 1 input: %s", lines[0])
	}
	if m := playerRE.FindStringSubmatch(lines[1]); len(m) == 2 {
		pos := must.Atoi(m[1])
		p2 = player{id: 2, pos: pos}
	} else {
		log.Fatalf("invalid player 2 input: %s", lines[0])
	}

	d := die{val: 1}
	rolls := 0
	for {
		d.roll(&p1)
		rolls += 3
		if p1.score >= 1000 {
			return p2.score * rolls
		}

		d.roll(&p2)
		rolls += 3
		if p2.score >= 1000 {
			return p1.score * rolls
		}
	}
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
