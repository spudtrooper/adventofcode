package lib

import (
	"log"
	"math"
	"strings"

	"github.com/spudtrooper/adventofcode/common/ints"
	"github.com/spudtrooper/adventofcode/common/must"
)

func Part1(input string) int {
	lines := must.ReadLines(input)
	tmpl := lines[0]
	rules := map[string]string{}
	for i := 1; i < len(lines); i++ {
		parts := strings.Split(lines[i], " -> ")
		if len(parts) == 2 {
			match, result := parts[0], parts[1]
			rules[match] = result
		}
	}

	step := func(tmpl string) string {
		chars := strings.Split(tmpl, "")
		var out []string
		for i := 0; i < len(chars)-1; i++ {
			ch := chars[i] + chars[i+1]
			matched := rules[ch]
			if matched == "" {
				log.Fatalf("no match for ch=%s for rules=%+v", ch, rules)
			}
			out = append(out, chars[i], matched)
		}
		out = append(out, chars[len(chars)-1])
		return strings.Join(out, "")
	}

	makeHist := func(s string) map[rune]int {
		hist := map[rune]int{}
		for _, c := range s {
			hist[c]++
		}
		return hist
	}

	log.Printf("Template:     %s", tmpl)
	s := tmpl
	for i := 0; i < 10; i++ {
		s = step(s)
		if len(s) > 200 {
			log.Printf("After step %d: (%d) ...", i+1, len(s))
		} else {
			log.Printf("After step %d: %s (%d)", i+1, s, len(s))
		}
	}

	hist := makeHist(s)
	min, max := math.MaxInt, math.MinInt
	for _, v := range hist {
		min = ints.Min(min, v)
		max = ints.Max(max, v)
	}

	res := max - min

	if want, got := res, part2(input, 10); want != got {
		log.Fatalf("comparing naive: want(%d) != got(%d)", want, got)
	}

	return res
}

func Part2(input string) int {
	return part2(input, 40)
}

func part2(input string, steps int) int {
	lines := must.ReadLines(input)
	tmpl := lines[0]
	rules := map[string]string{}
	for i := 1; i < len(lines); i++ {
		parts := strings.Split(lines[i], " -> ")
		if len(parts) == 2 {
			match, result := parts[0], parts[1]
			rules[match] = result
		}
	}

	type state struct {
		cnts map[string]int
		last string
	}

	makeHist := func(cnts state) map[string]int {
		hist := map[string]int{}
		for c, v := range cnts.cnts {
			hist[string(c[0])] += v
		}
		if cnts.last != "" {
			hist[cnts.last]++
		}
		return hist
	}

	initCounts := func(tmpl string) state {
		chars := strings.Split(tmpl, "")
		last := chars[len(chars)-1]
		cnts := state{cnts: map[string]int{}, last: last}
		for m := range rules {
			for _, v := range rules {
				cnts.cnts[m+v] = 0
			}
		}
		for i := range chars {
			if i < len(chars)-1 {
				k := chars[i] + chars[i+1]
				cnts.cnts[k]++
			}
		}
		return cnts
	}

	/*
		N       N       C       B
		N   C   N   B   C   H   B
		N B C C N B B B C B H C B
		NBBBCNCCNBBNBNBBCHBHHBCHB
	*/

	step := func(cnts state) state {
		new := state{cnts: map[string]int{}, last: cnts.last}
		for c, v := range cnts.cnts {
			if v == 0 {
				continue
			}
			matched := rules[c]
			if matched == "" {
				log.Fatalf("no match for ch=%s for rules=%+v", c, rules)
			}
			l, r := string(c[0])+matched, matched+string(c[1])
			new.cnts[l] += v
			new.cnts[r] += v
		}
		return new
	}

	cnts := initCounts(tmpl)
	for i := 0; i < steps; i++ {
		cnts = step(cnts)
	}

	hist := makeHist(cnts)
	min, max := math.MaxInt, math.MinInt
	for _, v := range hist {
		min = ints.Min(min, v)
		max = ints.Max(max, v)
	}

	return max - min
}
