package lib

import (
	"github.com/spudtrooper/adventofcode/common/ints"
	"github.com/spudtrooper/adventofcode/common/must"

	"github.com/flemeur/go-shortestpath/dijkstra"
)

func Part1(input string) int {
	b := ints.ReadBoardFromFile(input)

	return findShortestPath(b)
}

func findShortestPath(b ints.Board) int {
	width, height := b.Dims()

	var vs [][]*vertex
	for y := 0; y < height; y++ {
		var arr []*vertex
		for x := 0; x < width; x++ {
			arr = append(arr, makeVertex(b, x, y))
		}
		vs = append(vs, arr)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v := vs[y][x]
			if x > 0 {
				d := vs[y][x-1]
				v.edges = append(v.edges, edge{destination: d, weight: float64(b[d.y][d.x])})
			}
			if x < width-1 {
				d := vs[y][x+1]
				v.edges = append(v.edges, edge{destination: d, weight: float64(b[d.y][d.x])})
			}
			if y > 0 {
				d := vs[y-1][x]
				v.edges = append(v.edges, edge{destination: d, weight: float64(b[d.y][d.y])})
			}
			if y < height-1 {
				d := vs[y+1][x]
				v.edges = append(v.edges, edge{destination: d, weight: float64(b[d.y][d.x])})
			}
		}
	}

	start, end := vs[0][0], vs[height-1][width-1]
	path, err := dijkstra.ShortestPath(start, end)
	must.Check(err)

	var sum int
	for i := 1; i < len(path); i++ {
		v := path[i].(*vertex)
		sum += b[v.y][v.x]
	}

	return sum
}

type vertex struct {
	x, y  int
	edges []edge
}

func makeVertex(b ints.Board, x, y int) *vertex {
	return &vertex{x: x, y: y}
}

func (v *vertex) Edges() []dijkstra.Edge {
	edges := make([]dijkstra.Edge, len(v.edges))
	for i := range v.edges {
		edges[i] = v.edges[i]
	}
	return edges
}

// dijkstra.Edge implementation
type edge struct {
	destination *vertex
	weight      float64
}

func (e edge) Destination() dijkstra.Node {
	return e.destination
}

func (e edge) Weight() float64 {
	return e.weight
}

func Part2(input string) int {
	b := ints.ReadBoardFromFile(input)

	width, height := b.Dims()

	fullBoard := ints.Board{}
	for i := 0; i < 5*height; i++ {
		fullBoard = append(fullBoard, make([]int, width*5))
	}

	fullWidth, fullHeight := fullBoard.Dims()
	for y := 0; y < fullHeight; y++ {
		for x := 0; x < fullWidth; x++ {
			if x < width && y < height {
				fullBoard[y][x] = b[y][x]
			} else if x < width {
				fullBoard[y][x] = fullBoard[y-height][x]%9 + 1
			} else if y < height {
				fullBoard[y][x] = fullBoard[y][x-width]%9 + 1
			} else {
				fullBoard[y][x] = (fullBoard[y-height][x-width]%9+1)%9 + 1
			}
		}
	}

	return findShortestPath(fullBoard)
}
