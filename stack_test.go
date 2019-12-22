package main

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)

	if value, _ := s.Pop(); value != 2 {
		t.Errorf("got %d, want %d", value, 2)
	}

	if value, _ := s.Pop(); value != 1 {
		t.Errorf("got %d, want %d", value, 1)
	}

	if _, err := s.Pop(); err == nil {
		t.Errorf("got nil, want error")
	}
}
