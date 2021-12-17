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
			want:  6,
		},
		{
			name:  "testinput2",
			input: "testdata/testinput2.txt",
			want:  9,
		},
		{
			name:  "testinput3",
			input: "testdata/testinput3.txt",
			want:  14,
		},
		{
			name:  "testinput4",
			input: "testdata/testinput4.txt",
			want:  16,
		},
		{
			name:  "testinput5",
			input: "testdata/testinput5.txt",
			want:  12,
		},
		{
			name:  "testinput6",
			input: "testdata/testinput6.txt",
			want:  23,
		},
		{
			name:  "testinput7",
			input: "testdata/testinput7.txt",
			want:  31,
		},
		{
			name:  "part1",
			input: "testdata/input.txt",
			want:  871,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Part1(tc.input); want != got {
				t.Errorf("day16 Part1: want(%d) != got(%d)", want, got)
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
				t.Errorf("day16 Part2: want(%d) != got(%d)", want, got)
			}
		})
	}
}
