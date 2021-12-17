package geom

import "github.com/spudtrooper/adventofcode/common/ints"

func MakePoint(x, y int) Point {
	return &point{x: x, y: y}
}

func MakeRect(x1, x2, y1, y2 int) Rect {
	minX, maxX, minY, maxY := ints.Min(x1, x2), ints.Max(x1, x2), ints.Min(y1, y2), ints.Max(y1, y2)
	nw, se := MakePoint(minX, maxY), MakePoint(maxX, minY)
	return &rect{nw: nw, se: se}
}

func MakeRectFromPoints(nw, se Point) Rect {
	return &rect{nw: nw, se: se}
}

type Point interface {
	X() int
	Y() int
	Inside(r Rect) bool
	Move(dx, dy int) Point
	MoveBy(Point) Point
}

type Rect interface {
	NW() Point
	SE() Point
}

type point struct {
	x, y int
}

func (p *point) X() int { return p.x }
func (p *point) Y() int { return p.y }

func (p *point) Inside(r Rect) bool {
	if x, y := p.x, p.y; x >= r.NW().X() && x <= r.SE().X() && y <= r.NW().Y() && y >= r.SE().Y() {
		return true
	}
	return false
}

func (p *point) Move(dx, dy int) Point {
	return MakePoint(p.x+dx, p.y+dy)
}

func (p *point) MoveBy(o Point) Point {
	return p.Move(o.X(), o.Y())
}

type rect struct {
	nw, se Point
}

func (r *rect) NW() Point { return r.nw }
func (r *rect) SE() Point { return r.se }
