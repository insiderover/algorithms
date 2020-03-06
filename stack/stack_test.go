package stack_test

import (
	"algorithms/stack"
	"testing"
)

const stackSize = 10

func TestSimpleStack_Push(t *testing.T) {
	s := stack.New(stackSize)

	for i := 1; i <= stackSize; i++ {
		if err := s.Push(i); err != nil {
			t.Error(err)
		}
	}

	if err := s.Push(stackSize + 1); err == nil {
		t.Error("Expected error: stack is overflow, got nil")
	}
}

func TestSimpleStack_Pop(t *testing.T) {
	s := stack.New(stackSize)

	for i := 1; i <= stackSize; i++ {
		if err := s.Push(i); err != nil {
			t.Error(err)
		}
	}

	for i := 1; i <= stackSize; i++ {
		if _, err := s.Pop(); err != nil {
			t.Error(err)
		}
	}

	if _, err := s.Pop(); err == nil {
		t.Error("Expected error: stack is empty, got nil")
	}
}

func TestSimpleStack_Peek(t *testing.T) {
	s := stack.New(stackSize)

	if _, err := s.Peek(); err == nil {
		t.Error("Expected error: stack is empty, got nil")
	}

	for i := 1; i <= stackSize; i++ {
		if err := s.Push(i); err != nil {
			t.Error(err)
		}
	}

	if v, _ := s.Peek(); v != stackSize {
		t.Error("Expected value:", stackSize, "got", v)
	}
}
