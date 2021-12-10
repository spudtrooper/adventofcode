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

func (s *stack) Peek() (string, bool) {
	if s.Empty() {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	return element, true
}

var loud = false

func firstIllegalChar(chunk string) string {
	var s stack

	demand := func(want, right string, i int) (err bool) {
		if loud {
			log.Printf("%s:%d: demand want=%s right=%s stack=%v", chunk, i, want, right, s)
		}
		p, more := s.Pop()
		if !more {
			log.Fatalf("should have more pos:%d: %s", i, s)
		}
		if p != want {
			err = true
			log.Printf("Expected %s, but found %s instead at %d in %s", rightMatching[p], right, i, chunk)
		}
		return
	}

	for i, c := range strings.Split(chunk, "") {
		if loud {
			log.Printf("%s:%d: have c=%s stack=%v", chunk, i, c, s)
		}
		switch c {
		case "(", "[", "{", "<":
			if loud {
				log.Printf("%s:%d: pushing c=%s", chunk, i, c)
			}
			s.Push(c)
		case ")", "]", "}", ">":
			if err := demand(leftMatching[c], c, i); err {
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
		if char := firstIllegalChar(line); char != "" {
			score += points[char]
		}

	}

	return score
}

func firstIllegalCharWithStack(s *stack, chunk string) string {
	demand := func(want, right string, i int) (err bool) {
		if loud {
			log.Printf("%s:%d: demand want=%s right=%s stack=%v", chunk, i, want, right, s)
		}
		p, more := s.Pop()
		if !more {
			log.Fatalf("should have more pos:%d: %s", i, s)
		}
		if p != want {
			err = true
			log.Printf("Expected %s, but found %s instead at %d in %s", rightMatching[p], right, i, chunk)
		}
		return
	}

	for i, c := range strings.Split(chunk, "") {
		if loud {
			log.Printf("%s:%d: have c=%s stack=%v", chunk, i, c, s)
		}
		switch c {
		case "(", "[", "{", "<":
			if loud {
				log.Printf("%s:%d: pushing c=%s", chunk, i, c)
			}
			s.Push(c)
		case ")", "]", "}", ">":
			if err := demand(leftMatching[c], c, i); err {
				return c
			}
		default:
			log.Fatalf("Invalid character at pos %d: %s", i, chunk)
		}
	}
	return ""
}

func Part2(input string) int {
	var scores []int
	for _, line := range must.ReadStrings(input) {
		var s stack
		if char := firstIllegalCharWithStack(&s, line); char == "" {
			log.Printf("stack: %s", s)
			complete := ""
			var score int
			for {
				c, more := s.Pop()
				if !more {
					break
				}
				score *= 5
				score += autocompletePoints[rightMatching[c]]
				complete += rightMatching[c]
			}
			scores = append(scores, score)
			log.Printf("%s - %d total points.", complete, score)
		}

	}
	sort.Ints(scores)
	log.Printf("scores: %v", scores)
	middle := scores[len(scores)/2]
	return middle
}
