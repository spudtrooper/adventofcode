package main

import (
	"flag"
	"fmt"

	day15 "github.com/spudtrooper/adventofcode/2021/day15/lib"
)

var (
	input = flag.String("input", "2021/day15/lib/testdata/testinput.txt", "test input")
)

func main() {
	flag.Parse()
	fmt.Printf("day15 Part2: %d\n", day15.Part2(*input))
}
