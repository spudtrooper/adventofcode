package main

import (
	"flag"
	"fmt"

	day21 "github.com/spudtrooper/adventofcode/2021/day21/lib"
)

var (
	input  = flag.String("input", "2021/day21/lib/testdata/testinput.txt", "test input")
	packet = flag.String("packet", "", "input packet")
)

func main() {
	flag.Parse()
	var val int
	if *packet != "" {
		val = day21.Part2FromString(*packet)
	} else {
		val = day21.Part2(*input)
	}
	fmt.Printf("day21 Part2: %d\n", val)
}
