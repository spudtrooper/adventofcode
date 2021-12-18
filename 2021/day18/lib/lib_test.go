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
			want:  3488,
		},
		{
			name:  "part1",
			input: "testdata/input.txt",
			want:  4433,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Part1(tc.input); want != got {
				t.Errorf("day18 Part1: want(%d) != got(%d)", want, got)
			}
		})
	}
}

func Explode(n *node) *node {
	explode(n)
	return n
}

func TestPart1Explode(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "[[[[[9,8],1],2],3],4]",
			want:  "[[[[0,9],2],3],4]",
		},
		{
			input: "[7,[6,[5,[4,[3,2]]]]]",
			want:  "[7,[6,[5,[7,0]]]]",
		},
		{
			input: "[[6,[5,[4,[3,2]]]],1]",
			want:  "[[6,[5,[7,0]]],3]",
		},
		{
			input: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			want:  "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			input: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			want:  "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
	}
	for _, tc := range testCases {
		name := tc.input
		t.Run(name, func(t *testing.T) {
			if want, got := tc.want, Explode(Parse(tc.input)).String(); want != got {
				t.Errorf("day18 Explode: want(%s) != got(%s)", want, got)
			}
		})
	}
}

func TestPart1Reduce(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			want:  "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}
	for _, tc := range testCases {
		name := tc.input
		t.Run(name, func(t *testing.T) {
			if want, got := tc.want, Reduce(Parse(tc.input)).String(); want != got {
				t.Errorf("day18 Reduce: want(%s) != got(%s)", want, got)
			}
		})
	}
}

func TestPart1Magnitude(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			input: "[[1,2],[[3,4],5]]",
			want:  143,
		},
		{
			input: "[9,1]",
			want:  29,
		},
		{
			input: "[1,9]",
			want:  21,
		},
	}
	for _, tc := range testCases {
		name := tc.input
		t.Run(name, func(t *testing.T) {
			if want, got := tc.want, Magnitude(Parse(tc.input)); want != got {
				t.Errorf("day18 Magnitude: want(%d) != got(%d)", want, got)
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
				t.Errorf("day18 Part2: want(%d) != got(%d)", want, got)
			}
		})
	}
}
