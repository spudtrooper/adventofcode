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
			want:  10,
		},
		{
			name:  "testinput-larger",
			input: "testdata/testinput-larger.txt",
			want:  19,
		},
		{
			name:  "testinput-even-larger",
			input: "testdata/testinput-even-larger.txt",
			want:  226,
		},
		{
			name:  "part1",
			input: "testdata/input.txt",
			want:  3679,
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
			want:  36,
		},
		{
			name:  "testinput-larger",
			input: "testdata/testinput-larger.txt",
			want:  103,
		},
		{
			name:  "testinput-even-larger",
			input: "testdata/testinput-even-larger.txt",
			want:  3509,
		},
		{
			name:  "part2",
			input: "testdata/input.txt",
			want:  107395,
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
