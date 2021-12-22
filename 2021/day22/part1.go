package main

import (
	"flag"
	"fmt"

	day22 "github.com/spudtrooper/adventofcode/2021/day22/lib"
)

var (
	input  = flag.String("input", "2021/day22/lib/testdata/testinput.txt", "test input")
	packet = flag.String("packet", "", "input packet")
)

func main() {
	flag.Parse()
	var val int
	if *packet != "" {
		val = day22.Part1FromString(*packet)
	} else {
		val = day22.Part1(*input)
	}
	fmt.Printf("day22 Part1: %d\n", val)
}
