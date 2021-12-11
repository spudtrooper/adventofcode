package day06

import "testing"

const (
	testinput = "3,4,3,1,2"
	input     = "1,3,1,5,5,1,1,1,5,1,1,1,3,1,1,4,3,1,1,2,2,4,2,1,3,3,2,4,4,4,1,3,1,1,4,3,1,5,5,1,1,3,4,2,1,5,3,4,5,5,2,5,5,1,5,5,2,1,5,1,1,2,1,1,1,4,4,1,3,3,1,5,4,4,3,4,3,3,1,1,3,4,1,5,5,2,5,2,2,4,1,2,5,2,1,2,5,4,1,1,1,1,1,4,1,1,3,1,5,2,5,1,3,1,5,3,3,2,2,1,5,1,1,1,2,1,1,2,1,1,2,1,5,3,5,2,5,2,2,2,1,1,1,5,5,2,2,1,1,3,4,1,1,3,1,3,5,1,4,1,4,1,3,1,4,1,1,1,1,2,1,4,5,4,5,5,2,1,3,1,4,2,5,1,1,3,5,2,1,2,2,5,1,2,2,4,5,2,1,1,1,1,2,2,3,1,5,5,5,3,2,4,2,4,1,5,3,1,4,4,2,4,2,2,4,4,4,4,1,3,4,3,2,1,3,5,3,1,5,5,4,1,5,1,2,4,2,5,4,1,3,3,1,4,1,3,3,3,1,3,1,1,1,1,4,1,2,3,1,3,3,5,2,3,1,1,1,5,5,4,1,2,3,1,3,1,1,4,1,3,2,2,1,1,1,3,4,3,1,3"
)

func TestCount(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		days  int
		want  int
	}{
		{
			name:  "testinput-18",
			input: testinput,
			days:  18,
			want:  26,
		},
		{
			name:  "testinput-80",
			input: testinput,
			days:  80,
			want:  5934,
		},
		{
			name:  "part1",
			input: input,
			days:  80,
			want:  360268,
		},
		{
			name:  "testinput-256",
			input: testinput,
			days:  256,
			want:  26984457539,
		},
		{
			name:  "part2",
			input: input,
			days:  256,
			want:  1632146183902,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Count(tc.input, tc.days)
			if want, got := tc.want, res; want != got {
				t.Errorf("Count: want(%d) != got(%d)", want, got)
			}
		})
	}
}
