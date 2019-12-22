package main

import (
	"errors"
	"sync"
)

// ErrPopFromEmptyStack is the error resulting when trying to pop from an empty stack.
var ErrPopFromEmptyStack = errors.New("pop from empty stack")

// Stack is a thread-safe stack implementation of integers
type Stack struct {
	sync.Mutex
	values []int
	size   int
}

// NewStack creates a new empty stack of integers.
func NewStack() *Stack {
	return &Stack{
		values: []int{},
	}
}

// Push adds an element to the stack
func (s *Stack) Push(val int) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.values = append(s.values, val)
	s.size++
}

// Pop returns and removes the top of the stack, or an error if the stack is empty.
func (s *Stack) Pop() (int, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if s.size == 0 {
		return 0, ErrPopFromEmptyStack
	}

	val := s.values[s.size-1]

	s.values = s.values[:s.size-1]
	s.size--

	return val, nil
}
