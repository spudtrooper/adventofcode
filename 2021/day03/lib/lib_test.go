package day03

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "testinput",
			input: "testdata/testinput.txt",
			want:  198,
		},
		{
			name:  "part1",
			input: "testdata/input.txt",
			want:  2743844,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Part1(tc.input)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if want, got := tc.want, res.Multiply(); want != got {
				t.Errorf("Move: want(%d) != got(%d), res=%v", want, got, res)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "testinput",
			input: "testdata/testinput.txt",
			want:  230,
		},
		{
			name:  "part2",
			input: "testdata/input.txt",
			want:  6677951,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Part2(tc.input)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if want, got := tc.want, res.Multiply(); want != got {
				t.Errorf("Move: want(%d) != got(%d), res=%v", want, got, res)
			}
		})
	}
}
