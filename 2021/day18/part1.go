package main

import (
	"flag"
	"fmt"

	day18 "github.com/spudtrooper/adventofcode/2021/day18/lib"
)

var (
	input = flag.String("input", "2021/day18/lib/testdata/testinput.txt", "test input file")
	str   = flag.String("str", "", "test input string")
)

func main() {
	flag.Parse()
	// var val int
	// if *str != "" {
	// 	val = day18.Part1FromString(*str)
	// } else {
	// 	val = day18.Part1(*input)
	// }
	// fmt.Printf("day18 Part1: %d\n", val)

	fmt.Printf("Parse: %v\n", day18.Magnitude(day18.Parse(*str)))
}
