package stack

import (
	"errors"
)

type SimpleStack struct {
	size  int
	top   int
	stack []int
}

func New(size int) *SimpleStack {
	return &SimpleStack{size: size, top: -1, stack: make([]int, size, size)}
}

func (s *SimpleStack) Push(value int) error {
	if s.isFull() {
		return errors.New("stack is overflow")
	}

	s.top++
	s.stack[s.top] = value

	return nil
}

func (s *SimpleStack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}

	value := s.stack[s.top]
	s.top--

	return value, nil
}

func (s *SimpleStack) Peek() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.stack[s.top], nil
}

func (s *SimpleStack) isEmpty() bool {
	return s.top == -1
}

func (s *SimpleStack) isFull() bool {
	return s.size-1 == s.top
}
