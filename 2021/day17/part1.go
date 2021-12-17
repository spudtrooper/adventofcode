package main

import (
	"flag"
	"fmt"

	day17 "github.com/spudtrooper/adventofcode/2021/day17/lib"
)

var (
	input  = flag.String("input", "2021/day17/lib/testdata/testinput.txt", "test input")
	packet = flag.String("packet", "", "input packet")
)

func main() {
	flag.Parse()
	var val int
	if *packet != "" {
		val = day17.Part1FromString(*packet)
	} else {
		val = day17.Part1(*input)
	}
	fmt.Printf("day17 Part1: %d\n", val)
}
