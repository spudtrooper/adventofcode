package geom

func MakePoint3D(x, y, z int) Point3D {
	return &point3d{x: x, y: y, z: z}
}

type point3d struct {
	x, y, z int
}

func (p *point3d) X() int { return p.x }
func (p *point3d) Y() int { return p.y }
func (p *point3d) Z() int { return p.z }

func (p *point3d) Move(dx, dy, dz int) Point3D {
	return MakePoint3D(p.x+dx, p.y+dy, p.z+dz)
}

func (p *point3d) MoveBy(o Point3D) Point3D {
	return p.Move(o.X(), o.Y(), o.Z())
}

func (p *point3d) String() string        { return Point3DString(p) }
func (p *point3d) Equals(o Point3D) bool { return Point3DEquals(p, o) }
