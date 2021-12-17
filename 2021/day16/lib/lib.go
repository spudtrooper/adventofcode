package lib

import (
	"log"
	"math"

	"github.com/spudtrooper/adventofcode/common/ints"
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
)

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

type packets struct {
	packets []packet
}

func parseInt(s string) int {
	return int(must.ParseInt(s, 2, 64))
}

func versions(p packet) int {
	if p.id == 4 {
		return p.version
	}
	res := p.version
	ps := p.payload.(*packets)
	for _, p := range ps.packets {
		res += versions(p)
	}
	return res
}

func parse(s *scanner) packet {
	version := parseInt(s.next(3))
	id := parseInt(s.next(3))

	if id == 4 {
		for num := ""; s.hasMore(); {
			next := s.next(5)
			fst, lst := next[0:1], next[1:]
			num += lst
			if fst == "0" {
				return packet{
					version: version,
					id:      id,
					payload: &literalValue{num: parseInt(num)},
				}
			}
		}
	}

	ps := []packet{}
	if i := parseInt(s.next(1)); i == 0 {
		l := parseInt(s.next(15))
		rest := s.next(l)
		for sub := makeScanner(rest); sub.hasMore(); {
			ps = append(ps, parse(sub))
		}
	} else {
		l := parseInt(s.next(11))
		for i := 0; i < l; i++ {
			p := parse(s)
			ps = append(ps, p)
		}
	}

	return packet{
		version: version,
		id:      id,
		payload: &packets{
			packets: ps,
		},
	}
}

func value(p packet) int {
	switch p.id {
	case 4:
		p := p.payload.(*literalValue)
		return p.num
	case 0:
		ps := p.payload.(*packets)
		res := 0
		for _, p := range ps.packets {
			res += value(p)
		}
		return res
	case 1:
		ps := p.payload.(*packets)
		res := 1
		for _, p := range ps.packets {
			res *= value(p)
		}
		return res
	case 2:
		ps := p.payload.(*packets)
		res := math.MaxInt
		for _, p := range ps.packets {
			res = ints.Min(res, value(p))
		}
		return res
	case 3:
		ps := p.payload.(*packets)
		res := math.MinInt
		for _, p := range ps.packets {
			res = ints.Max(res, value(p))
		}
		return res
	case 5:
		ps := p.payload.(*packets)
		if fst, snd := value(ps.packets[0]), value(ps.packets[1]); fst > snd {
			return 1
		}
		return 0
	case 6:
		ps := p.payload.(*packets)
		if fst, snd := value(ps.packets[0]), value(ps.packets[1]); fst < snd {
			return 1
		}
		return 0
	case 7:
		ps := p.payload.(*packets)
		if fst, snd := value(ps.packets[0]), value(ps.packets[1]); fst == snd {
			return 1
		}
		return 0
	}

	panic("fuck my life")
}

func Part1(input string) int {
	packet := must.ReadAllFile(input)

	var bin string
	for _, r := range packet {
		bin += h2b[string(r)]
	}

	s := makeScanner(bin)
	p := parse(s)
	return versions(p)
}

func Part2(input string) int {
	packet := must.ReadAllFile(input)
	return Part2FromString(packet)
}

func Part2FromString(packet string) int {
	var bin string
	for _, r := range packet {
		bin += h2b[string(r)]
	}

	s := makeScanner(bin)
	p := parse(s)

	return value(p)
}
