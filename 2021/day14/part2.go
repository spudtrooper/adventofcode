package main

import (
	"flag"
	"fmt"

	day14 "github.com/spudtrooper/adventofcode/2021/day14/lib"
)

var (
	input = flag.String("input", "2021/day14/lib/testdata/testinput.txt", "test input")
)

func main() {
	flag.Parse()
	fmt.Printf("day14 Part2: %d\n", day14.Part2(*input))
}
