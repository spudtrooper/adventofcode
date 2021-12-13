package main

import (
	"flag"
	"fmt"

	day13 "github.com/spudtrooper/adventofcode/2021/day13/lib"
)

var (
	input = flag.String("input", "2021/day13/lib/testdata/testinput.txt", "test input")
)

func main() {
	flag.Parse()
	fmt.Printf("day13 Part1: %d\n", day13.Part1(*input))
}
