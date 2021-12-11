package day06

import (
	"strings"

	"github.com/spudtrooper/adventofcode/common/must"
)

func Count(input string, days int) int {
	var fs []int
	for _, s := range strings.Split(input, ",") {
		fs = append(fs, must.Atoi(s))
	}

	type state map[int]int // [day]count

	s := state{}
	for i := 1; i < 9; i++ {
		s[i] = 0
	}
	for _, f := range fs {
		s[int(f)]++
	}
	for day := 0; day < days; day++ {
		zeros := s[0]
		for i := 1; i < 9; i++ {
			s[i-1] = s[i]
		}
		s[6] += zeros
		s[8] = zeros
	}
	var res int
	for _, c := range s {
		res += c
	}
	return res
}
