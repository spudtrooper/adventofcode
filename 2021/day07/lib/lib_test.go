package day07

import "testing"

const (
	testFn = false
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
			want:  37,
		},
		{
			name:  "part1",
			input: "testdata/input.txt",
			want:  328187,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Part1(tc.input); want != got {
				t.Errorf("Part1: want(%d) != got(%d)", want, got)
			}
			if testFn {
				if want, got := tc.want, Part1Func(tc.input); want != got {
					t.Errorf("Part1Func: want(%d) != got(%d)", want, got)
				}
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
			want:  168,
		},
		{
			name:  "part2",
			input: "testdata/input.txt",
			want:  91257582,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Part2(tc.input); want != got {
				t.Errorf("Part2: want(%d) != got(%d)", want, got)
			}
			if testFn {
				if want, got := tc.want, Part2Func(tc.input); want != got {
					t.Errorf("Part2Func: want(%d) != got(%d)", want, got)
				}
			}
		})
	}
}
