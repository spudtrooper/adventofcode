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
			want:  35,
		},
		{
			name:  "part1",
			input: "testdata/input.txt",
			want:  5361,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Part1(tc.input); want != got {
				t.Errorf("day20 Part1: want(%d) != got(%d)", want, got)
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
			want:  -1, // TODO
		},
		// {
		// 	name:  "part2",
		//	input: "testdata/input.txt",
		//	want:  -1, // TODO
		// },
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Part2(tc.input); want != got {
				t.Errorf("day20 Part2: want(%d) != got(%d)", want, got)
			}
		})
	}
}
