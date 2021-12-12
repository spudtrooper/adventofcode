package main

import (
	"flag"
	"fmt"

	day12 "github.com/spudtrooper/adventofcode/2021/day12/lib"
)

var (
	input = flag.String("input", "2021/day12/testdata/testinput.txt", "test input")
)

func main() {
	flag.Parse()
	fmt.Printf("day12 Part2: %d\n", day12.Part2(*input))
}