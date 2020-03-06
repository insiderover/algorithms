package stack_test

import (
	"algorithms/stack"
	"testing"
)

const (
	testStackSize = 10
)

func TestSimpleStack_Push(t *testing.T) {
	s := getEmptyStack(testStackSize)

	for i := 1; i <= testStackSize; i++ {
		if err := s.Push(i); err != nil {
			t.Error("Expected error is nil, got", err)
		}
	}

	if err := s.Push(testStackSize + 1); err == nil {
		t.Error("Expected error: stack is overflow, got nil")
	}
}

func TestSimpleStack_Pop(t *testing.T) {
	s := getFillStack(testStackSize)

	for i := 1; i <= testStackSize; i++ {
		if _, err := s.Pop(); err != nil {
			t.Error(err)
		}
	}

	if _, err := s.Pop(); err == nil {
		t.Error("Expected error: stack is empty, got nil")
	}
}

func TestSimpleStack_Peek(t *testing.T) {
	s := getEmptyStack(testStackSize)

	if _, err := s.Peek(); err == nil {
		t.Error("Expected error: stack is empty, got nil")
	}

	s = getFillStack(testStackSize)

	if v, _ := s.Peek(); v != testStackSize {
		t.Error("Expected value:", testStackSize, "got", v)
	}
}

func getEmptyStack(size int) *stack.SimpleStack {
	return stack.New(size)
}

func getFillStack(size int) *stack.SimpleStack {
	s := stack.New(size)

	for i := 1; i <= testStackSize; i++ {
		_ = s.Push(i)
	}

	return s
}
