package day10

import (
	"log"
	"sort"
	"strings"

	"github.com/spudtrooper/adventofcode/common/must"
)

var (
	matching = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
)

// https://www.educative.io/edpresso/how-to-implement-a-stack-in-golang
type stack []string

func (s *stack) Empty() bool {
	return len(*s) == 0
}

func (s *stack) Push(str string) {
	*s = append(*s, str)
}

func (s *stack) Pop() string {
	if s.Empty() {
		return ""
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

type stackFinder struct{ stack }

func (s *stackFinder) FirstIllegalChar(chunk string) string {
	for i, c := range strings.Split(chunk, "") {
		switch c {
		case "(", "[", "{", "<":
			s.Push(c)
		case ")", "]", "}", ">":
			if p := s.Pop(); matching[p] != c {
				log.Printf("Expected %s, but found %s instead at %d in %s", matching[p], c, i, chunk)
				return c
			}
		default:
			log.Fatalf("Invalid character at pos %d: %s", i, chunk)
		}
	}
	return ""
}

func Part1(input string) int {
	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	var score int
	for _, line := range must.ReadStrings(input) {
		var s stackFinder
		if c := s.FirstIllegalChar(line); c != "" {
			score += points[c]
		}
	}
	return score
}

func Part2(input string) int {
	points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	var scores []int
	for _, line := range must.ReadStrings(input) {
		var s stackFinder
		if c := s.FirstIllegalChar(line); c == "" {
			var score int
			for {
				p := s.Pop()
				if p == "" {
					break
				}
				score *= 5
				score += points[matching[p]]
			}
			scores = append(scores, score)
		}

	}
	sort.Ints(scores)
	log.Printf("scores: %v", scores)
	return scores[len(scores)/2]
}
