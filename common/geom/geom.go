package geom

import "fmt"

type Point interface {
	X() int
	Y() int
	Inside(r Rect) bool
	Move(dx, dy int) Point
	MoveBy(Point) Point
	Equals(o Point) bool
	String() string
}

func PointString(p Point) string {
	return fmt.Sprintf("(%d,%d)", p.X(), p.Y())
}

func PointEquals(p, o Point) bool {
	return p.X() == o.X() && p.Y() == o.Y()
}

type Point3D interface {
	X() int
	Y() int
	Z() int
	Move(dx, dy, dz int) Point3D
	MoveBy(Point3D) Point3D
	Equals(o Point3D) bool
}

func Point3DString(p Point3D) string {
	return fmt.Sprintf("(%d,%d,%d)", p.X(), p.Y(), p.Z())
}

func Point3DEquals(p, o Point3D) bool {
	return p.X() == o.X() && p.Y() == o.Y() && p.Z() == o.Z()
}

type Rect interface {
	NW() Point
	SE() Point
}
