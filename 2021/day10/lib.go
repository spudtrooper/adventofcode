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

// https://www.educative.io/edpresso/how-to-implement-a-stack-in-golang
type stack []string

func (s *stack) Empty() bool {
	return len(*s) == 0
}

func (s *stack) Push(str string) {
	*s = append(*s, str)
}

func (s *stack) Pop() (string, bool) {
	if s.Empty() {
		return "", true
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, false
}

type stackFinder struct{ stack }

func (s *stackFinder) FirstIllegalChar(chunk string) string {
	for i, c := range strings.Split(chunk, "") {
		switch c {
		case "(", "[", "{", "<":
			s.Push(c)
		case ")", "]", "}", ">":
			p, done := s.Pop()
			if done {
				log.Fatalf("should have more pos:%d: %s", i, s)
			}
			if want := leftMatching[c]; p != want {
				log.Printf("Expected %s, but found %s instead at %d in %s", rightMatching[p], c, i, chunk)
				return c
			}
		default:
			log.Fatalf("Invalid character at pos %d: %s", i, chunk)
		}
	}
	return ""
}

func Part1(input string) int {
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
	var scores []int
	for _, line := range must.ReadStrings(input) {
		var s stackFinder
		if c := s.FirstIllegalChar(line); c == "" {
			complete := ""
			var score int
			for {
				p, done := s.Pop()
				if done {
					break
				}
				score *= 5
				score += autocompletePoints[rightMatching[p]]
				complete += rightMatching[p]
			}
			log.Printf("%s - %d total points.", complete, score)
			scores = append(scores, score)
		}

	}
	sort.Ints(scores)
	log.Printf("scores: %v", scores)
	return scores[len(scores)/2]
}
