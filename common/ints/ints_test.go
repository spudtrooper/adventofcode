package ints

import (
	"log"
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "empty",
			input: []int{},
			want:  0,
		},
		{
			name:  "one",
			input: []int{1},
			want:  1,
		},
		{
			name:  "1-2",
			input: []int{1, 2},
			want:  1,
		},
		{
			name:  "2-1",
			input: []int{2, 1},
			want:  1,
		},
		{
			name:  "many",
			input: []int{2, 1, 3, 5, -7, 4},
			want:  -7,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Min(FromArray(&tc.input)); want != got {
				t.Errorf("Min: want(%d) != got(%d)", want, got)
			}
			if want, got := tc.want, FromArray(&tc.input).Min(); want != got {
				t.Errorf("Min: want(%d) != got(%d)", want, got)
			}
		})
	}
}

func TestRange2(t *testing.T) {
	r := FromRange(1, 3)
	if !r.HasNext() {
		t.Errorf("should have next after 0")
	}
	if want, got := 1, r.Next(); want != got {
		t.Errorf("Next: want(%d) != got(%d)", want, got)
	}
	if !r.HasNext() {
		t.Errorf("should have next after 1")
	}
	if want, got := 2, r.Next(); want != got {
		t.Errorf("Next: want(%d) != got(%d)", want, got)
	}
	if !r.HasNext() {
		t.Errorf("should have next after 2")
	}
	if want, got := 3, r.Next(); want != got {
		t.Errorf("Next: want(%d) != got(%d)", want, got)
	}
	if r.HasNext() {
		t.Errorf("should not have next after 3")
	}
}

func TestRange3(t *testing.T) {
	r := FromRange(1, 1)
	if !r.HasNext() {
		t.Errorf("should have next after 0")
	}
	if want, got := 1, r.Next(); want != got {
		t.Errorf("Next: want(%d) != got(%d)", want, got)
	}
	if r.HasNext() {
		t.Errorf("should not have next after 1")
	}
}

func TestRangeInteg(t *testing.T) {
	for r := FromRange(1, 1000); ; {
		if !r.HasNext() {
			break
		}
		x := r.Next()
		for r := FromRange(1, 1000); ; {
			if !r.HasNext() {
				break
			}
			y := r.Next()
			log.Printf("x.y: %d:%d", x, y)
		}
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "empty",
			input: []int{},
			want:  0,
		},
		{
			name:  "one",
			input: []int{1},
			want:  1,
		},
		{
			name:  "1-2",
			input: []int{1, 2},
			want:  2,
		},
		{
			name:  "2-1",
			input: []int{2, 1},
			want:  2,
		},
		{
			name:  "many",
			input: []int{2, 1, 3, 5, -7, 4},
			want:  5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Max(FromArray(&tc.input)); want != got {
				t.Errorf("Max: want(%d) != got(%d)", want, got)
			}
			if want, got := tc.want, FromArray(&tc.input).Max(); want != got {
				t.Errorf("Max: want(%d) != got(%d)", want, got)
			}
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "empty",
			input: []int{},
			want:  0,
		},
		{
			name:  "one",
			input: []int{1},
			want:  1,
		},
		{
			name:  "1-2",
			input: []int{1, 2},
			want:  3,
		},
		{
			name:  "2-1",
			input: []int{2, 1},
			want:  3,
		},
		{
			name:  "many",
			input: []int{2, 1, 3, 5, -7, 4},
			want:  8,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Sum(FromArray(&tc.input)); want != got {
				t.Errorf("Sum: want(%d) != got(%d)", want, got)
			}
			if want, got := tc.want, FromArray(&tc.input).Sum(); !reflect.DeepEqual(want, got) {
				t.Errorf("Reduce: want(%d) != got(%d)", want, got)
			}

		})
	}
}

func TestMap(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "empty",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "one",
			input: []int{1},
			want:  []int{2},
		},
		{
			name:  "1-2",
			input: []int{1, 2},
			want:  []int{2, 4},
		},
		{
			name:  "2-1",
			input: []int{2, 1},
			want:  []int{4, 2},
		},
		{
			name:  "many",
			input: []int{2, 1, 3, 5, -7, 4},
			want:  []int{4, 2, 6, 10, -14, 8},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Map(FromArray(&tc.input), func(x int) int { return 2 * x }).Vec().Array(); !reflect.DeepEqual(want, got) {
				t.Errorf("Map: want(%d) != got(%d)", want, got)
			}
			if want, got := tc.want, FromArray(&tc.input).Map(func(x int) int { return 2 * x }).Vec().Array(); !reflect.DeepEqual(want, got) {
				t.Errorf("Map: want(%d) != got(%d)", want, got)
			}
		})
	}
}

func TestRange(t *testing.T) {
	testCases := []struct {
		name     string
		from, to int
		want     []int
	}{
		{
			name: "same",
			from: 3,
			to:   3,
			want: []int{3},
		},
		{
			name: "1-2",
			from: 1,
			to:   2,
			want: []int{1, 2},
		},
		{
			name: "many",
			from: 1,
			to:   5,
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Range(tc.from, tc.to); !reflect.DeepEqual(want, got) {
				t.Errorf("Range: want(%d) != got(%d)", want, got)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		base  int
		want  int
	}{
		{
			name:  "empty",
			input: []int{},
			want:  0,
		},
		{
			name:  "one",
			input: []int{1},
			want:  1,
		},
		{
			name:  "1-2",
			input: []int{1, 2},
			want:  3,
		},
		{
			name:  "2-1",
			input: []int{2, 1},
			want:  3,
		},
		{
			name:  "many",
			input: []int{2, 1, 3, 5, -7, 4},
			want:  8,
		},
		{
			name:  "many base=3",
			input: []int{2, 1, 3, 5, -7, 4},
			base:  3,
			want:  11,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, Reduce(FromArray(&tc.input), tc.base, func(a, b int) int { return a + b }); want != got {
				t.Errorf("Reduce: want(%d) != got(%d)", want, got)
			}
			if want, got := tc.want, FromArray(&tc.input).Reduce(tc.base, func(a, b int) int { return a + b }); want != got {
				t.Errorf("Reduce: want(%d) != got(%d)", want, got)
			}
		})
	}
}
