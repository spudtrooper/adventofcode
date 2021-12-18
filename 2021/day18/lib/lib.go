package lib

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/spudtrooper/adventofcode/common"
	"github.com/spudtrooper/adventofcode/common/must"
)

type node struct {
	num  int
	l, r *node
}

func (n *node) String() string {
	if n.l != nil {
		return fmt.Sprintf("[%v,%v]", n.l, n.r)
	}
	return fmt.Sprintf("%d", n.num)
}

type nodes []*node

func (ns nodes) IndexOf(n *node) int {
	for i, t := range ns {
		if t == n {
			return i
		}
	}
	return -1
}

func (n *node) ToList() nodes {
	if n.l != nil {
		var res []*node
		res = append(res, n.l.ToList()...)
		res = append(res, n.r.ToList()...)
		return res
	}
	return []*node{n}
}

func parseNode(s common.Scanner, tab string) *node {
	ch := s.Next(1)
	if ch == "[" {
		l := parseNode(s, tab+" ")
		if ch := s.Next(1); ch != "," {
			log.Fatalf(`expecting  "," instead found %s`, ch)
		}
		r := parseNode(s, tab+" ")
		if ch := s.Next(1); ch != "]" {
			log.Fatalf(`expecting  "]" instead found %s`, ch)
		}
		return &node{l: l, r: r}
	}
	n := ch
	isDigit := func(s string) bool {
		if _, err := strconv.Atoi(s); err == nil {
			return true
		}
		return false
	}
	for isDigit(s.Peek()) {
		n += s.Next(1)
	}
	num := must.Atoi(n)
	return &node{num: num}
}

func Parse(input string) *node {
	s := common.MakeScanner(input)
	return parseNode(s, " ")
}

func Add(a, b *node) *node {
	return add(a, b)
}

func add(a, b *node) *node {
	return &node{l: a, r: b}
}

func findNodeToExplode(n *node, depth int) *node {
	if depth >= 4 && n.l != nil && n.l.l == nil && n.r != nil && n.r.l == nil {
		return n
	}
	if n.l != nil {
		if ex := findNodeToExplode(n.l, depth+1); ex != nil {
			return ex
		}
	}
	if n.r != nil {
		if ex := findNodeToExplode(n.r, depth+1); ex != nil {
			return ex
		}
	}
	return nil
}

func Explode(n *node) *node {
	explode(n)
	return n
}

func explode(in *node) bool {
	if ex := findNodeToExplode(in, 0); ex != nil {
		ns := in.ToList()
		if i := ns.IndexOf(ex.l); i > 0 {
			ns[i-1].num += ex.l.num
		}
		if i := ns.IndexOf(ex.r); i < len(ns)-1 {
			ns[i+1].num += ex.r.num
		}
		*ex = node{}
		return true
	}
	return false
}

func split(in *node) bool {
	for _, n := range in.ToList() {
		if n.num < 10 {
			continue
		}
		l, r := n.num/2, n.num/2+n.num%2
		*n = node{l: &node{num: l}, r: &node{num: r}}
		return true
	}
	return false
}

func Reduce(n *node) *node {
	for {
		if explode(n) || split(n) {
			continue
		}
		break
	}
	return n
}

func Magnitude(n *node) int {
	return mag(n)
}

func mag(n *node) int {
	if n.l != nil {
		return 3*mag(n.l) + 2*mag(n.r)
	}
	return n.num
}

func Part1FromString(input string) int {
	lines := strings.Split(input, "\n")
	l := Parse(lines[0])
	for i := 1; i < len(lines); i++ {
		r := Parse(lines[i])
		n := Add(l, r)
		l = Reduce(n)
	}

	return Magnitude(l)
}

func Part2FromString(input string) int {
	lines := strings.Split(input, "\n")
	max := math.MinInt
	for i, ls := range lines {
		for j, rs := range lines {
			if i == j {
				continue
			}
			l, r := Parse(ls), Parse(rs)
			n := Add(l, r)
			red := Reduce(n)
			if m := Magnitude(red); m > max {
				max = m
			}
		}
	}
	return max
}

func Part1(input string) int {
	return Part1FromString(must.ReadAllFile(input))
}

func Part2(input string) int {
	return Part2FromString(must.ReadAllFile(input))
}
