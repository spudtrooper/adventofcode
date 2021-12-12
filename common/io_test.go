package common

import (
	"reflect"
	"strconv"
	"testing"
)

func TestReadFile(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		parser func(s string) (interface{}, error)
		want   []interface{}
	}{
		{
			name:   "empty",
			input:  "testdata/empty.txt",
			parser: nil,
			want:   []interface{}{},
		},
		{
			name:   "ints",
			input:  "testdata/ints.txt",
			parser: func(s string) (interface{}, error) { return strconv.Atoi(s) },
			want:   []interface{}{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arr, err := ReadFile(tc.input, tc.parser)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if want, got := tc.want, arr; !reflect.DeepEqual(want, got) {
				t.Errorf("ReadFile: want(%v) != got(%v)", want, got)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		sep    string
		parser func(s string) (interface{}, error)
		want   []interface{}
	}{
		{
			name:   "empty",
			input:  "",
			parser: nil,
			want:   []interface{}{},
		},
		{
			name:   "ints",
			input:  "1,2,3",
			sep:    ",",
			parser: func(s string) (interface{}, error) { return strconv.Atoi(s) },
			want:   []interface{}{1, 2, 3},
		},
		{
			name:   "ints-|",
			input:  "1|2|3",
			sep:    "|",
			parser: func(s string) (interface{}, error) { return strconv.Atoi(s) },
			want:   []interface{}{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arr, err := Split(tc.input, tc.sep, tc.parser)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if want, got := tc.want, arr; !reflect.DeepEqual(want, got) {
				t.Errorf("ReadFile: want(%v) != got(%v)", want, got)
			}
		})
	}
}

func TestReadStrings(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "empty",
			input: "testdata/empty.txt",
			want:  []string{},
		},
		{
			name:  "ints",
			input: "testdata/ints.txt",
			want:  []string{"1", "2", "3"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arr, err := ReadLines(tc.input)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if want, got := tc.want, arr; !reflect.DeepEqual(want, got) {
				t.Errorf("ReadFile: want(%v) != got(%v)", want, got)
			}
		})
	}
}
