package day10

import "testing"

// func TestPart1(t *testing.T) {
// 	testCases := []struct {
// 		name  string
// 		input string
// 		want  int
// 	}{
// 		{
// 			name:  "testinput-small",
// 			input: "testdata/testinput-small.txt",
// 			want:  1197,
// 		},
// 		{
// 			name:  "testinput-small2",
// 			input: "testdata/testinput-small2.txt",
// 			want:  1197,
// 		},
// 		{
// 			name:  "testinput",
// 			input: "testdata/testinput.txt",
// 			want:  26397,
// 		},
// 		{
// 			name:  "part1",
// 			input: "testdata/input.txt",
// 			want:  167379,
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			if want, got := tc.want, Part1(tc.input); want != got {
// 				t.Errorf("Part1: want(%d) != got(%d)", want, got)
// 			}
// 		})
// 	}
// }

func TestPart2(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "testinput",
			input: "testdata/testinput.txt",
			want:  -2, // TODO
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
				t.Errorf("Part2: want(%d) != got(%d)", want, got)
			}
		})
	}
}
