package main

import (
	"flag"
	"fmt"

	day18 "github.com/spudtrooper/adventofcode/2021/day18/lib"
)

var (
	input  = flag.String("input", "2021/day18/lib/testdata/testinput.txt", "test input")
	packet = flag.String("packet", "", "input packet")
)

func main() {
	flag.Parse()
	var val int
	if *packet != "" {
		val = day18.Part2FromString(*packet)
	} else {
		val = day18.Part2(*input)
	}
	fmt.Printf("day18 Part2: %d\n", val)
}
