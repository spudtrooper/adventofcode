package day10

import (
	"log"
	"sort"
	"strings"

	"github.com/spudtrooper/adventofcode/common"
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

type stackFinder struct{ common.Stack }

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
	for _, line := range must.ReadLines(input) {
		s := &stackFinder{common.MakeStack()}
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
	for _, line := range must.ReadLines(input) {
		s := &stackFinder{common.MakeStack()}
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
