package main

import (
	"flag"
	"fmt"

	day16 "github.com/spudtrooper/adventofcode/2021/day16/lib"
)

var (
	input  = flag.String("input", "2021/day16/lib/testdata/testinput.txt", "test input")
	packet = flag.String("packet", "", "input packet")
)

func main() {
	flag.Parse()
	var val int
	if *packet != "" {
		val = day16.Part2FromString(*packet)
	} else {
		val = day16.Part2(*input)
	}
	fmt.Printf("day16 Part2: %d\n", val)
}
