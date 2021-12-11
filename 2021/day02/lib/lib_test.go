package day02

import "testing"

func TestSubMove(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "testinput",
			input: "testdata/testinput.txt",
			want:  150,
		},
		{
			name:  "part1",
			input: "testdata/input.txt",
			want:  2039256,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var s sub
			mul, err := s.Move(tc.input)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if want, got := tc.want, mul; want != got {
				t.Errorf("Move: want(%d) != got(%d)", want, got)
			}
		})
	}
}

func TestSub2Move(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "testinput",
			input: "testdata/testinput.txt",
			want:  900,
		},
		{
			name:  "part2",
			input: "testdata/input.txt",
			want:  1856459736,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var s sub2
			mul, err := s.Move(tc.input)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if want, got := tc.want, mul; want != got {
				t.Errorf("Move: want(%d) != got(%d)", want, got)
			}
		})
	}
}
