package day10

import (
	"log"
	"sort"
	"strings"

	"github.com/spudtrooper/adventofcode/common/must"
)

var (
	points = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	autocompletePoints = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	leftMatching = map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}
	rightMatching = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
)

type stack []string

func (s *stack) Empty() bool {
	return len(*s) == 0
}

func (s *stack) Push(str string) {
	*s = append(*s, str)
}

func (s *stack) Pop() (string, bool) {
	if s.Empty() {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func firstIllegalChar(s *stack, chunk string) string {
	// demand := func(want, right string, i int) (err bool) {
	// 	p, more := s.Pop()
	// 	if !more {
	// 		log.Fatalf("should have more pos:%d: %s", i, s)
	// 	}
	// 	if p != want {
	// 		err = true
	// 		log.Printf("Expected %s, but found %s instead at %d in %s", rightMatching[p], right, i, chunk)
	// 	}
	// 	return
	// }

	for i, c := range strings.Split(chunk, "") {
		switch c {
		case "(", "[", "{", "<":
			s.Push(c)
		case ")", "]", "}", ">":
			p, more := s.Pop()
			if !more {
				log.Fatalf("should have more pos:%d: %s", i, s)
			}
			if want := leftMatching[c]; p != want {
				log.Printf("Expected %s, but found %s instead at %d in %s", rightMatching[p], c, i, chunk)
				return c
			}
			// if err := demand(leftMatching[c], c, i); err {
			// 	return c
			// }
		default:
			log.Fatalf("Invalid character at pos %d: %s", i, chunk)
		}
	}
	return ""
}

func Part1(input string) int {
	var score int
	for _, line := range must.ReadStrings(input) {
		var s stack
		if c := firstIllegalChar(&s, line); c != "" {
			score += points[c]
		}
	}
	return score
}

func Part2(input string) int {
	var scores []int
	for _, line := range must.ReadStrings(input) {
		var s stack
		if c := firstIllegalChar(&s, line); c == "" {
			complete := ""
			var score int
			for {
				p, more := s.Pop()
				if !more {
					break
				}
				score *= 5
				score += autocompletePoints[rightMatching[p]]
				complete += rightMatching[p]
			}
			scores = append(scores, score)
			log.Printf("%s - %d total points.", complete, score)
		}

	}
	sort.Ints(scores)
	log.Printf("scores: %v", scores)
	return scores[len(scores)/2]
}
