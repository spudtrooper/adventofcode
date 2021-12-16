package lib

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
			want:  40,
		},
		{
			name:  "part1",
			input: "testdata/input.txt",
			want:  702,
		},
		{
			name:  "part1",
			input: "testdata/testinput-large.txt",
			want:  315,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Part1(tc.input); want != got {
				t.Errorf("day15 Part1: want(%d) != got(%d)", want, got)
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
			want:  320,
		},
		{
			name:  "part2",
			input: "testdata/input.txt",
			want:  2956,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Part2(tc.input); want != got {
				t.Errorf("day15 Part2: want(%d) != got(%d)", want, got)
			}
		})
	}
}
