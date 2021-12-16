package main

import (
	"flag"
	"fmt"

	day16 "github.com/spudtrooper/adventofcode/2021/day16/lib"
)

var (
	input = flag.String("input", "2021/day16/lib/testdata/testinput.txt", "test input")
)

func main() {
	flag.Parse()
	fmt.Printf("day16 Part1: %d\n", day16.Part1(*input))
}
