package common

import (
	"reflect"
	"testing"
)

func TestScsannerNext(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		size  int
		want  []string
	}{
		{
			name:  "empty",
			input: "",
			size:  1,
			want:  []string{},
		},
		{
			input: "123",
			size:  1,
			want:  []string{"1", "2", "3"},
		},
		{
			input: "123456",
			size:  3,
			want:  []string{"123", "456"},
		},
		{
			input: "12345",
			size:  3,
			want:  []string{"123", "45"},
		},
	}
	for _, tc := range testCases {
		name := tc.name
		if name == "" {
			name = tc.input
		}
		t.Run(name, func(t *testing.T) {
			s := MakeScanner(tc.input)
			have := []string{}
			for {
				if !s.HasMore() {
					break
				}
				t := s.Next(tc.size)
				have = append(have, t)
			}
			if want, got := tc.want, have; !reflect.DeepEqual(want, got) {
				t.Errorf("Next: want(%v) != got(%v)", want, got)
			}
		})
	}
}

func TestScsannerPeek(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		size  int
		want  string
	}{
		{
			name:  "empty",
			input: "",
			size:  1,
			want:  "",
		},
		{
			input: "123",
			size:  1,
			want:  "1",
		},
		{
			input: "123456",
			size:  3,
			want:  "123",
		},
		{
			input: "12",
			size:  3,
			want:  "12",
		},
	}
	for _, tc := range testCases {
		name := tc.name
		if name == "" {
			name = tc.input
		}
		t.Run(name, func(t *testing.T) {
			s := MakeScanner(tc.input)
			if want, got := tc.want, s.Peek(tc.size); !reflect.DeepEqual(want, got) {
				t.Errorf("Peek: want(%v) != got(%v)", want, got)
			}
		})
	}
}
