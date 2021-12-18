package common

// https://www.educative.io/edpresso/how-to-implement-a-stack-in-golang

type stack []string

type Stack interface {
	Empty() bool
	Push(str string)
	Pop() string
}

func MakeStack() Stack {
	return &stack{}
}

func (s *stack) Empty() bool {
	return len(*s) == 0
}

func (s *stack) Push(str string) {
	*s = append(*s, str)
}

func (s *stack) Pop() string {
	if s.Empty() {
		return ""
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}
