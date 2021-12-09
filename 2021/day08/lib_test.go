package day08

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
			want:  26,
		},
		{
			name:  "part1",
			input: "testdata/input.txt",
			want:  375,
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
			name:  "testinput-small",
			input: "testdata/testinput-small.txt",
			want:  5353,
		},
		{
			name:  "testinput",
			input: "testdata/testinput.txt",
			want:  61229,
		},
		{
			name:  "part2",
			input: "testdata/input.txt",
			want:  1019355,
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
