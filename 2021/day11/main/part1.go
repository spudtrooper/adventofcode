package main

import (
	"flag"
	"fmt"

	day11 "github.com/spudtrooper/adventofcode/2021/day11/lib"
)

var (
	input = flag.String("input", "2021/day11/testdata/testinput.txt", "test input")
)

func main() {
	flag.Parse()
	fmt.Printf("day11 Part1: %d\n", day11.Part1(*input))
}
