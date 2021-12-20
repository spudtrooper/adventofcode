package geom

func MakePoint(x, y int) Point {
	return &point{x: x, y: y}
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
