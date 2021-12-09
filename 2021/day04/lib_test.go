package day04

import "testing"

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "testinput",
			input: "testdata/testinput.txt",
			want:  4512,
		},
		{
			name:  "testinput2",
			input: "testdata/testinput2.txt",
			want:  242,
		},
		{
			name:  "part1",
			input: "testdata/input.txt",
			want:  45031,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Part1(tc.input); want != got {
				t.Errorf("Part1: want(%d) != got(%d)", want, got)
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
			want:  1924,
		},
		{
			name:  "part2",
			input: "testdata/input.txt",
			want:  2568,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Part2(tc.input); want != got {
				t.Errorf("Part2: want(%d) != got(%d)", want, got)
			}
		})
	}
}
