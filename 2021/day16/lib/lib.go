package lib

import (
	"log"

	"github.com/spudtrooper/adventofcode/common/must"
)

var (
	h2b = map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
	b2h = map[string]string{
		"0000": "0",
		"0001": "1",
		"0010": "2",
		"0011": "3",
		"0100": "4",
		"0101": "5",
		"0110": "6",
		"0111": "7",
		"1000": "8",
		"1001": "9",
		"1010": "A",
		"1011": "B",
		"1100": "C",
		"1101": "D",
		"1110": "E",
		"1111": "F",
	}
)

// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

type scanner struct {
	input string
	cur   int
}

func makeScanner(s string) *scanner {
	return &scanner{input: s}
}

func (s *scanner) next(n int) string {
	res := s.input[s.cur : s.cur+n]
	s.cur += n
	log.Printf("after next(%d)=%s", n, res)
	return res
}

func (s *scanner) hasMore() bool {
	return s.cur < len(s.input)
}

type packet struct {
	version int
	id      int
	payload interface{}
}

type literalValue struct {
	num int
}

func parseInt(s string) int {
	return int(must.ParseInt(s, 2, 64))
}

func parse(s *scanner) []*packet {
	version := parseInt(s.next(3))
	id := parseInt(s.next(3))

	p := &packet{
		version: version,
		id:      id,
	}

	if id == 4 {
		num := ""
		for s.hasMore() {
			next := s.next(5)
			fst, lst := next[0:1], next[1:]
			num += lst
			if fst == "0" {
				break
			}
		}
		p.payload = &literalValue{
			num: parseInt(num),
		}
		return []*packet{p}
	}

	res := []*packet{p}

	i := parseInt(s.next(1))
	if i == 0 {
		l := parseInt(s.next(15))
		packets := s.next(l)
		for sub := makeScanner(packets); sub.hasMore(); {
			res = append(res, parse(sub)...)
		}
	} else {
		l := parseInt(s.next(11))
		for i := 0; i < l; i++ {
			p := parse(s)
			res = append(res, p...)
		}
	}

	return res
}

func Part1(input string) int {
	packet := must.ReadAllFile(input)

	var bin string
	for _, r := range packet {
		bin += h2b[string(r)]
	}

	s := makeScanner(bin)
	ps := parse(s)

	res := 0
	for _, p := range ps {
		log.Printf("p: %+v", p)
		res += p.version
	}
	return res
}

func Part2(input string) int {
	for _, line := range must.ReadLines(input) {
		// TODO
		if false {
			log.Println(line)
		}
	}
	return -1
}
