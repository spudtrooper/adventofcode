package geom

type Point interface {
	X() int
	Y() int
	Inside(r Rect) bool
	Move(dx, dy int) Point
	MoveBy(Point) Point
}

type Point3D interface {
	X() int
	Y() int
	Z() int
	Move(dx, dy, dz int) Point3D
	MoveBy(Point3D) Point3D
}

type Rect interface {
	NW() Point
	SE() Point
}
