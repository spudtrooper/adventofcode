package common

type BoolBoard3D interface {
	Dims() (width, height, depth int)
	Get(y, x, z int) bool
	Set(y, x, z int, s bool)
	Traverse(func(y, x, z int, s bool))
}

func MakeBoolBoard3D(width, height, depth int, ctor func(y, x, z int) bool) BoolBoard3D {
	var b boolBoard3D
	for y := 0; y < height; y++ {
		var rows [][]bool
		for x := 0; x < width; x++ {
			var row []bool
			for z := 0; z < depth; z++ {
				c := ctor(y, x, z)
				row = append(row, c)
			}
			rows = append(rows, row)
		}
		b = append(b, rows)
	}
	return b
}

func MakeEmptyBoolBoard3D(width, height, depth int) BoolBoard3D {
	var b boolBoard3D = make([][][]bool, height)
	for y := 0; y < height; y++ {
		b[y] = make([][]bool, width)
		for x := 0; x < width; x++ {
			b[y][x] = make([]bool, depth)
		}
	}
	return b
}

func BoolBoard3DIdentFn(v bool) func(y, x, z int) bool {
	return func(y, x, z int) bool {
		return v
	}
}

func (b boolBoard3D) Traverse(f func(y, x, z int, s bool)) {
	for y, rows := range b {
		for x, row := range rows {
			for z, v := range row {
				f(y, x, z, v)
			}
		}
	}
}

type boolBoard3D [][][]bool

func (b boolBoard3D) Dims() (width, height, depth int) {
	height, width, depth = len(b), len(b[0]), len(b[0][0])
	return
}

func (b boolBoard3D) Get(y, x, z int) bool {
	return b[y][x][z]
}

func (b boolBoard3D) Set(y, x, z int, v bool) {
	b[y][x][z] = v
}

// func (b boolboard3D) String() string {
// 	var lines []string
// 	for _, row := range b {
// 		var bs []string
// 		for _, b := range row {
// 			v := "F"
// 			if b {
// 				v = "T"
// 			}
// 			bs = append(bs, v)
// 		}
// 		lines = append(lines, strings.Join(bs, ""))
// 	}
// 	return strings.Join(lines, "\n")
// }
