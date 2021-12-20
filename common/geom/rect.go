package geom

import "github.com/spudtrooper/adventofcode/common/ints"

func MakeRect(x1, x2, y1, y2 int) Rect {
	minX, maxX, minY, maxY := ints.Min(x1, x2), ints.Max(x1, x2), ints.Min(y1, y2), ints.Max(y1, y2)
	nw, se := MakePoint(minX, maxY), MakePoint(maxX, minY)
	return &rect{nw: nw, se: se}
}

func MakeRectFromPoints(nw, se Point) Rect {
	return &rect{nw: nw, se: se}
}

type rect struct {
	nw, se Point
}

func (r *rect) NW() Point { return r.nw }
func (r *rect) SE() Point { return r.se }
