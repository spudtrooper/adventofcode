package lib

import (
	"log"
	"strconv"
	"strings"

	"github.com/spudtrooper/adventofcode/common/must"
)

type instr struct {
	op   string
	a, b string
}

type alu struct {
	w, y, x, z int
}

type stream struct {
	cur int
	buf []int
}

func (i *stream) next() int {
	res := i.buf[i.cur]
	i.cur++
	return res
}

func execute(instrs []instr, alu *alu, in stream) {

	put := func(s int, reg string) {
		switch reg {
		case "w":
			alu.w = s
		case "x":
			alu.x = s
		case "y":
			alu.y = s
		case "z":
			alu.z = s
		default:
			log.Fatalf("put: invalid reg: %s", reg)
		}
	}

	get := func(reg string) int {
		if v, err := strconv.Atoi(reg); err == nil {
			return v
		}
		switch reg {
		case "w":
			return alu.w
		case "x":
			return alu.x
		case "y":
			return alu.y
		case "z":
			return alu.z
		default:
			log.Fatalf("get: invalid reg: %s", reg)
		}
		return -1
	}

	for _, instr := range instrs {
		switch op, a, b := instr.op, instr.a, instr.b; op {
		case "inp":
			v := in.next()
			put(v, a)
		case "add":
			x, y := get(a), get(b)
			v := x + y
			put(v, a)
		case "mul":
			x, y := get(a), get(b)
			v := x * y
			put(v, a)
		case "div":
			x, y := get(a), get(b)
			v := 0
			if y != 0 {
				v = x / y
			}
			put(v, a)
		case "mod":
			x, y := get(a), get(b)
			v := 0
			if x >= 0 && y > 0 {
				v = x % y
			}
			put(v, a)
		case "eql":
			x, y := get(a), get(b)
			v := 0
			if x == y {
				v = 1
			}
			put(v, a)
		default:
			log.Fatalf("invalid op: %s", op)
		}
	}
}

func Part1FromString(input string) int {
	var instrs []instr
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		op, a, b := parts[0], parts[1], ""
		if len(parts) == 3 {
			b = parts[2]
		}
		instr := instr{op, a, b}
		instrs = append(instrs, instr)
	}

	for i := 0; i < 10; i++ {
		log.Printf("--------- %d", i)
		in := stream{buf: []int{i}}
		alu := alu{}
		execute(instrs, &alu, in)
		log.Printf("alu: %+v", alu)
	}

	return -1
}

func Part2FromString(input string) int {
	for _, line := range strings.Split(input, "\n") {
		// TODO
		if false {
			log.Println(line)
		}
	}
	return -1
}

func Part1(input string) int {
	return Part1FromString(must.ReadAllFile(input))
}

func Part2(input string) int {
	return Part2FromString(must.ReadAllFile(input))
}
