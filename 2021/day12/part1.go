package main

import (
	"flag"
	"fmt"

	day12 "github.com/spudtrooper/adventofcode/2021/day12/lib"
)

var (
	input = flag.String("input", "2021/day12/lib/testdata/testinput.txt", "test input")
)

func main() {
	flag.Parse()
	fmt.Printf("day12 Part1: %d\n", day12.Part1(*input))
}
