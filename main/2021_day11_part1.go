package main

import (
	"flag"
	"fmt"

	"github.com/spudtrooper/adventofcode/2021/day11"
)

var (
	input = flag.String("input", "2021/day11/testdata/testinput.txt", "test input")
)

func main() {
	flag.Parse()
	fmt.Printf("Part1: %d\n", day11.Part1(*input))
}
