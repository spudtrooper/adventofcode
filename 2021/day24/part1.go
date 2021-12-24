package main

import (
	"flag"
	"fmt"

	day24 "github.com/spudtrooper/adventofcode/2021/day24/lib"
)

var (
	input  = flag.String("input", "2021/day24/lib/testdata/testinput.txt", "test input")
	packet = flag.String("packet", "", "input packet")
)

func main() {
	flag.Parse()
	var val int
	if *packet != "" {
		val = day24.Part1FromString(*packet)
	} else {
		val = day24.Part1(*input)
	}
	fmt.Printf("day24 Part1: %d\n", val)
}
