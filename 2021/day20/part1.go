package main

import (
	"flag"
	"fmt"

	day20 "github.com/spudtrooper/adventofcode/2021/day20/lib"
)

var (
	input  = flag.String("input", "2021/day20/lib/testdata/testinput.txt", "test input")
	packet = flag.String("packet", "", "input packet")
)

func main() {
	flag.Parse()
	var val int
	if *packet != "" {
		val = day20.Part1FromString(*packet)
	} else {
		val = day20.Part1(*input)
	}
	fmt.Printf("day20 Part1: %d\n", val)
}
