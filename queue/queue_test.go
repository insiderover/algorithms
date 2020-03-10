package queue

import (
	"testing"
)

const (
	testQueueSize = 10
)

func TestQueue_Insert(t *testing.T) {
	queue := NewQueue(testQueueSize)

	for i := 1; i <= testQueueSize; i++ {
		if err := queue.Insert(i); err != nil {
			t.Error("Expected error is nil, got", err)
		}
	}

	if err := queue.Insert(testQueueSize + 1); err == nil {
		t.Error("Expected error: queue is full, got nil")
	}

	// for better coverage
	_, _ = queue.Remove()
	_ = queue.Insert(20)
}

func TestQueue_Remove(t *testing.T) {
	queue := NewQueue(testQueueSize)
	fillQueue(queue)

	for i := 1; i <= testQueueSize; i++ {
		if _, err := queue.Remove(); err != nil {
			t.Error(err)
		}
	}

	if _, err := queue.Remove(); err == nil {
		t.Error("Expected error: queue is empty, got nil")
	}
}

func TestQueue_Size(t *testing.T) {
	queue := NewQueue(testQueueSize)
	fillQueue(queue)

	if size := queue.Size(); size != testQueueSize {
		t.Errorf("Expected size: %d, got %d", testQueueSize, size)
	}
}

func TestPriorityQueue_Insert(t *testing.T) {
	queue := NewPriorityQueue(testQueueSize)

	for i := 1; i <= testQueueSize; i++ {
		if err := queue.Insert(i); err != nil {
			t.Error("Expected error is nil, got", err)
		}
	}

	if err := queue.Insert(testQueueSize + 1); err == nil {
		t.Error("Expected error: queue is full, got nil")
	}

	// for better coverage
	_, _ = queue.Remove()
	_ = queue.Insert(0)
}

func TestPriorityQueue_Remove(t *testing.T) {
	queue := NewPriorityQueue(testQueueSize)
	fillQueue(queue)

	for i := 1; i <= testQueueSize; i++ {
		v, err := queue.Remove()
		if err != nil {
			t.Error(err)
		}
		if v != i {
			t.Errorf("Expected %d, got %d", i, v)
		}
	}

	if _, err := queue.Remove(); err == nil {
		t.Error("Expected error: queue is empty, got nil")
	}
}

func TestPriorityQueue_Size(t *testing.T) {
	queue := NewPriorityQueue(testQueueSize)
	fillQueue(queue)

	if size := queue.Size(); size != testQueueSize {
		t.Errorf("Expected size: %d, got %d", testQueueSize, size)
	}
}

func fillQueue(queue IQueue) {
	for i := 1; i <= testQueueSize; i++ {
		_ = queue.Insert(i)
	}
}
