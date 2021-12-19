package common

import "github.com/fatih/color"

type Scanner interface {
	Next(int) string
	Peek(int) string
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
	res := s.peek(n)
	s.cur += n
	return res
}

func (s *scanner) Peek(n int) string {
	return s.peek(n)
}

func (s *scanner) peek(n int) string {
	end := s.cur + n
	if end > len(s.input) {
		end = len(s.input)
	}
	return s.input[s.cur:end]
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
