package day01

import (
	"testing"
)

func TestCountLevels(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "testinput",
			input: "testdata/testinput.txt",
			want:  7,
		},
		{
			name:  "part1",
			input: "testdata/input.txt",
			want:  1451,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := CountLevels(tc.input)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if want, got := tc.want, output.Increased(); want != got {
				t.Errorf("CountLevels: want(%d) != got(%d), output=%s", want, got, output)
			}
		})
	}
}

func TestCountLevelsGrouped(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		clusterSize int
		want        int
	}{
		{
			name:        "testinput-1",
			input:       "testdata/testinput.txt",
			clusterSize: 1,
			want:        7,
		},
		{
			name:        "input-1",
			input:       "testdata/input.txt",
			clusterSize: 1,
			want:        1451,
		},
		{
			name:        "testinput-3",
			input:       "testdata/testinput.txt",
			clusterSize: 3,
			want:        5,
		},
		{
			name:        "part2",
			input:       "testdata/input.txt",
			clusterSize: 3,
			want:        1395,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := CountLevelsGrouped(tc.input, tc.clusterSize)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if want, got := tc.want, output.Increased(); want != got {
				t.Errorf("CountLevelsGrouped: want(%d) != got(%d), output=%s", want, got, output)
			}
		})
	}
}
