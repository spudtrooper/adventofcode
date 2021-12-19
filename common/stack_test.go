package common

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := MakeStack()
	if want, got := true, s.Empty(); want != got {
		t.Errorf("Empty: want(%v) != got(%v)", want, got)
	}
	if want, got := "", s.Pop(); want != got {
		t.Errorf("Pop: want(%v) != got(%v)", want, got)
	}
	s.Push("one")
	s.Push("two")
	s.Push("three")
	if want, got := false, s.Empty(); want != got {
		t.Errorf("Empty: want(%v) != got(%v)", want, got)
	}
	if want, got := "three", s.Pop(); want != got {
		t.Errorf("Pop: want(%v) != got(%v)", want, got)
	}
	if want, got := false, s.Empty(); want != got {
		t.Errorf("Empty: want(%v) != got(%v)", want, got)
	}
	if want, got := "two", s.Pop(); want != got {
		t.Errorf("Pop: want(%v) != got(%v)", want, got)
	}
	s.Push("four")
	if want, got := false, s.Empty(); want != got {
		t.Errorf("Empty: want(%v) != got(%v)", want, got)
	}
	if want, got := "four", s.Pop(); want != got {
		t.Errorf("Pop: want(%v) != got(%v)", want, got)
	}
	if want, got := false, s.Empty(); want != got {
		t.Errorf("Empty: want(%v) != got(%v)", want, got)
	}
	if want, got := "one", s.Pop(); want != got {
		t.Errorf("Pop: want(%v) != got(%v)", want, got)
	}
	if want, got := true, s.Empty(); want != got {
		t.Errorf("Empty: want(%v) != got(%v)", want, got)
	}
	if want, got := "", s.Pop(); want != got {
		t.Errorf("Pop: want(%v) != got(%v)", want, got)
	}
}
