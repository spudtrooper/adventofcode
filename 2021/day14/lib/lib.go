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

	return max - min
}

func Part2(input string) int {
	for _, line := range must.ReadLines(input) {
		// TODO
		if false {
			log.Println(line)
		}
	}
	return -1
}
