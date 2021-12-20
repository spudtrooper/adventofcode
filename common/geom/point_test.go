package geom

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPointInside(t *testing.T) {
	type testCase struct {
		name  string
		point Point
		rect  Rect
		want  bool
	}
	testCases := []testCase{
		{
			name:  "trivial",
			point: MakePoint(0, 0),
			rect:  MakeRectFromPoints(MakePoint(0, 0), MakePoint(0, 0)),
			want:  true,
		},
		{
			name:  "left",
			point: MakePoint(-6, 0),
			rect:  MakeRectFromPoints(MakePoint(-5, 5), MakePoint(5, -5)),
			want:  false,
		},
		{
			name:  "above",
			point: MakePoint(0, 6),
			rect:  MakeRectFromPoints(MakePoint(-5, 5), MakePoint(5, -5)),
			want:  false,
		},
		{
			name:  "right",
			point: MakePoint(6, 0),
			rect:  MakeRectFromPoints(MakePoint(-5, 5), MakePoint(5, -5)),
			want:  false,
		},
		{
			name:  "below",
			point: MakePoint(0, -6),
			rect:  MakeRectFromPoints(MakePoint(-5, 5), MakePoint(5, -5)),
			want:  false,
		},
		{
			name:  "inside",
			point: MakePoint(0, 0),
			rect:  MakeRectFromPoints(MakePoint(-5, 5), MakePoint(5, -5)),
			want:  true,
		},
	}
	for x := -5; x <= 5; x++ {
		for y := -5; y <= 5; y++ {
			testCases = append(testCases, testCase{
				name:  fmt.Sprintf("x=%d y=%d", x, y),
				point: MakePoint(x, y),
				rect:  MakeRectFromPoints(MakePoint(-5, 5), MakePoint(5, -5)),
				want:  true,
			})
		}
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, tc.point.Inside(tc.rect); want != got {
				t.Errorf("Inside: want(%t) != got(%t) point=%v rect=%v", want, got, tc.point, tc.rect)
			}
		})
	}
}

func TestPointMove(t *testing.T) {
	testCases := []struct {
		name   string
		point  Point
		dx, dy int
		want   Point
	}{
		{
			name:  "no move",
			point: MakePoint(0, 0),
			dx:    0,
			dy:    0,
			want:  MakePoint(0, 0),
		},
		{
			name:  "move left",
			point: MakePoint(5, 5),
			dx:    -1,
			dy:    0,
			want:  MakePoint(4, 5),
		},
		{
			name:  "move right",
			point: MakePoint(5, 5),
			dx:    1,
			dy:    0,
			want:  MakePoint(6, 5),
		},
		{
			name:  "move up",
			point: MakePoint(5, 5),
			dx:    0,
			dy:    -1,
			want:  MakePoint(5, 4),
		},
		{
			name:  "move down",
			point: MakePoint(5, 5),
			dx:    0,
			dy:    1,
			want:  MakePoint(5, 6),
		},
		{
			name:  "move up/left",
			point: MakePoint(5, 5),
			dx:    -1,
			dy:    1,
			want:  MakePoint(4, 6),
		},
		{
			name:  "move up/right",
			point: MakePoint(5, 5),
			dx:    1,
			dy:    -1,
			want:  MakePoint(6, 4),
		},
		{
			name:  "move down/left",
			point: MakePoint(5, 5),
			dx:    -1,
			dy:    -1,
			want:  MakePoint(4, 4),
		},
		{
			name:  "move down/right",
			point: MakePoint(5, 5),
			dx:    1,
			dy:    -1,
			want:  MakePoint(6, 4),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, tc.point.Move(tc.dx, tc.dy); !reflect.DeepEqual(want, got) {
				t.Errorf("Move(%d,%d): want(%v) != got(%v) point=%v", tc.dx, tc.dy, want, got, tc.point)
			}
		})
	}
}

func TestPointMoveBy(t *testing.T) {
	testCases := []struct {
		name  string
		point Point
		by    Point
		want  Point
	}{
		{
			name:  "no move",
			point: MakePoint(0, 0),
			by:    MakePoint(0, 0),
			want:  MakePoint(0, 0),
		},
		{
			name:  "move left",
			point: MakePoint(5, 5),
			by:    MakePoint(-1, 0),
			want:  MakePoint(4, 5),
		},
		{
			name:  "move right",
			point: MakePoint(5, 5),
			by:    MakePoint(1, 0),
			want:  MakePoint(6, 5),
		},
		{
			name:  "move up",
			point: MakePoint(5, 5),
			by:    MakePoint(0, -1),
			want:  MakePoint(5, 4),
		},
		{
			name:  "move down",
			point: MakePoint(5, 5),
			by:    MakePoint(0, 1),
			want:  MakePoint(5, 6),
		},
		{
			name:  "move up/left",
			point: MakePoint(5, 5),
			by:    MakePoint(-1, 1),
			want:  MakePoint(4, 6),
		},
		{
			name:  "move up/right",
			point: MakePoint(5, 5),
			by:    MakePoint(1, -1),
			want:  MakePoint(6, 4),
		},
		{
			name:  "move down/left",
			point: MakePoint(5, 5),
			by:    MakePoint(-1, -1),
			want:  MakePoint(4, 4),
		},
		{
			name:  "move down/right",
			point: MakePoint(5, 5),
			by:    MakePoint(1, -1),
			want:  MakePoint(6, 4),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, tc.point.MoveBy(tc.by); !reflect.DeepEqual(want, got) {
				t.Errorf("Move(%v): want(%v) != got(%v) point=%v", tc.by, want, got, tc.point)
			}
		})
	}
}
