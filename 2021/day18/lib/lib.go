package lib

import (
	"fmt"
	"log"
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

// [[[[[9,8],1],2],3],4]
func parseNode(s common.Scanner, tab string) *node {
	if ch := s.Peek(); ch == "[" {
		s.Next(1)
		a := parseNode(s, tab+" ")
		if ch := s.Next(1); ch != "," {
			log.Fatalf(`expecting  "," instead found %s`, ch)
		}
		b := parseNode(s, tab+" ")
		if ch := s.Next(1); ch != "]" {
			log.Fatalf(`expecting  "]" instead found %s`, ch)
		}
		return &node{l: a, r: b}
	}
	n := s.Next(1)
	for {
		if !isDigit(s.Peek()) {
			break
		}
		n += s.Next(1)
	}
	num := must.Atoi(n)
	return &node{num: num}
}

func isDigit(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func Parse(input string) *node {
	s := common.MakeScanner(input)
	return parseNode(s, " ")
}

func add(a, b *node) *node {
	return &node{l: a, r: b}
}

func findNodeToExplode(n *node, depth int) *node {
	// log.Printf("findNodeToExplode: node:%v depth:%d", n, depth)
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

func explode(n *node) bool {
	if ex := findNodeToExplode(n, 0); ex != nil {
		ns := n.ToList()
		li, ri := ns.IndexOf(ex.l), ns.IndexOf(ex.r)
		if li > 0 {
			ns[li-1].num += ex.l.num
		}
		if ri < len(ns)-1 {
			ns[ri+1].num += ex.r.num
		}
		ex.l = nil
		ex.r = nil
		ex.num = 0
		return true
	}
	return false
}

func split(n *node) bool {
	for _, s := range n.ToList() {
		if s.num >= 10 {
			l, r := s.num/2, s.num/2+s.num%2
			s.l = &node{num: l}
			s.r = &node{num: r}
			s.num = 0
			return true
		}
	}
	return false
}

func Reduce(n *node) *node {
	log.Printf("reduce          : %v", n)
	for {
		if !reduce(n) {
			break
		}
	}
	return n
}

func reduce(n *node) bool {
	if explode(n) {
		log.Printf("after explode   : %v", n)
		return true
	}
	if split(n) {
		log.Printf("after split     : %v", n)
		return true
	}
	return false
}

func Magnitude(n *node) int {
	return magnitude(n)
}

func magnitude(n *node) int {
	if n.l != nil {
		return 3*magnitude(n.l) + 2*magnitude(n.r)
	}
	return n.num
}

func Part1FromString(input string) int {
	lines := strings.Split(input, "\n")
	l := Parse(lines[0])
	for i := 1; i < len(lines); i++ {
		r := Parse(lines[i])
		n := add(l, r)
		l = Reduce(n)
	}

	log.Printf("result: %v", l)

	return Magnitude(l)
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
