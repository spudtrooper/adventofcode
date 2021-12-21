package lib

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/spudtrooper/adventofcode/common/must"
	"github.com/thomaso-mirodin/intmath/intgr"
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
	r1, r2, r3 := d.next(), d.next(), d.next()
	p.pos = (p.pos + r1 + r2 + r3) % 10
	p.score += p.pos + 1
	log.Printf("Player %d rolls %d+%d+%d and moves to space %d for a total score of %d", p.id, r1, r2, r3, p.pos, p.score)
}

func Part1FromString(input string) int {
	lines := strings.Split(input, "\n")
	var p1, p2 player
	if m := playerRE.FindStringSubmatch(lines[0]); len(m) == 2 {
		pos := must.Atoi(m[1])
		p1 = player{id: 1, pos: pos - 1}
	} else {
		log.Fatalf("invalid player 1 input: %s", lines[0])
	}
	if m := playerRE.FindStringSubmatch(lines[1]); len(m) == 2 {
		pos := must.Atoi(m[1])
		p2 = player{id: 2, pos: pos - 1}
	} else {
		log.Fatalf("invalid player 2 input: %s", lines[1])
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
	lines := strings.Split(input, "\n")
	var p1, p2 player
	if m := playerRE.FindStringSubmatch(lines[0]); len(m) == 2 {
		pos := must.Atoi(m[1])
		p1 = player{id: 1, pos: pos - 1}
	} else {
		log.Fatalf("invalid player 1 input: %s", lines[0])
	}
	if m := playerRE.FindStringSubmatch(lines[1]); len(m) == 2 {
		pos := must.Atoi(m[1])
		p2 = player{id: 2, pos: pos - 1}
	} else {
		log.Fatalf("invalid player 2 input: %s", lines[1])
	}

	log.Printf("initial: %+v and %+v", p1, p2)

	type state struct {
		p1, p2 player
	}

	freqs := map[int]int{
		3: 1, // 111
		4: 3, // 112, 121, 211
		5: 6, // 113, 122, 212, 221, 131, 311
		6: 7, // ...
		7: 6,
		8: 3,
		9: 1,
	}

	// TODO: Learn to map from objects correctly
	states := map[string]state{}
	counts := map[string]int{}
	initState := state{
		p1: p1,
		p2: p2,
	}
	states[fmt.Sprintf("%+v", initState)] = initState
	counts[fmt.Sprintf("%+v", initState)] = 1

	var wins1, wins2 int
	for len(states) > 0 {
		newStates := map[string]state{}
		newCounts := map[string]int{}
		for k, s := range states {
			cnt := counts[k]
			for roll1, freq1 := range freqs {
				cnt1 := cnt * freq1
				s1 := s
				s1.p1.pos = (s1.p1.pos + roll1) % 10
				s1.p1.score += s1.p1.pos + 1
				if s1.p1.score >= 21 {
					wins1 += cnt1
					continue
				}
				for roll2, freq2 := range freqs {
					cnt2 := cnt1 * freq2
					s2 := s1
					s2.p2.pos = (s2.p2.pos + roll2) % 10
					s2.p2.score += s2.p2.pos + 1
					if s2.p2.score >= 21 {
						wins2 += cnt2
						continue
					}
					newStates[fmt.Sprintf("%+v", s2)] = s2
					newCounts[fmt.Sprintf("%+v", s2)] += cnt2
				}
			}
		}
		states = newStates
		counts = newCounts
	}

	log.Printf("wins1: %d", wins1)
	log.Printf("wins2: %d", wins2)

	return intgr.Max(wins1, wins2)
}

func Part1(input string) int {
	return Part1FromString(must.ReadAllFile(input))
}

func Part2(input string) int {
	return Part2FromString(must.ReadAllFile(input))
}
