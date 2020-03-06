package stack_test

import (
	"algorithms/stack"
	"testing"
)

const (
	testBenchStackSize = 1_000_000
)

func BenchmarkSimpleStack_Push(b *testing.B) {
	b.ReportAllocs()

	var testStack *stack.SimpleStack
	testStack = getEmptyStack(testBenchStackSize)

	for j := 1; j <= testStackSize; j++ {
		_ = testStack.Push(j)
	}
}

func BenchmarkSimpleStack_Pop(b *testing.B) {
	b.ReportAllocs()

	var testStack *stack.SimpleStack
	testStack = getFillStack(testBenchStackSize)

	for j := 1; j <= testStackSize; j++ {
		_, _ = testStack.Pop()
	}
}

func BenchmarkSimpleStack_Peek(b *testing.B) {
	b.ReportAllocs()

	var testStack *stack.SimpleStack
	testStack = getFillStack(testBenchStackSize)

	for j := 1; j <= testStackSize; j++ {
		_, _ = testStack.Peek()
	}
}
