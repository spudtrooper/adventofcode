package common

import "github.com/fatih/color"

type Scanner interface {
	Next(int) string
	Peek() string
	DebugString() string
	HasMore() bool
}

type scanner struct {
	input string
	cur   int
}

func MakeScanner(s string) *scanner {
	return &scanner{input: s}
}

func (s *scanner) Next(n int) string {
	res := s.input[s.cur : s.cur+n]
	s.cur += n
	return res
}

func (s *scanner) Peek() string {
	return s.input[s.cur : s.cur+1]
}

func (s *scanner) Pos() int {
	return s.cur
}

func (s *scanner) DebugString() string {
	var (
		colorPast = color.New(color.FgGreen)
		colorCur  = color.New(color.FgHiWhite).Add(color.Bold).Add(color.Underline).Add(color.Italic) // a little much
	)
	var res string
	for i, c := range s.input {
		ch := string(c)
		if i < s.cur {
			ch = colorPast.Sprintf("%s", ch)
		} else if i == s.cur {
			ch = colorCur.Sprintf("%s", ch)
		}
		res += ch
	}
	return res
}

func (s *scanner) HasMore() bool {
	return s.cur < len(s.input)
}
