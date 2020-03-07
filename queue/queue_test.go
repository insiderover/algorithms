package queue

import "testing"

const (
	testQueueSize = 10
)

func TestQueue_Insert(t *testing.T) {
	queue := getEmptyQueue(testQueueSize)

	for i := 1; i <= testQueueSize; i++ {
		if err := queue.Insert(i); err != nil {
			t.Error("Expected error is nil, got", err)
		}
	}

	if err := queue.Insert(testQueueSize + 1); err == nil {
		t.Error("Expected error: queue is full, got nil")
	}
}

func TestQueue_Remove(t *testing.T) {
	queue := getFillQueue(testQueueSize)

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
	queue := getFillQueue(testQueueSize)

	if size := queue.Size(); size != testQueueSize {
		t.Errorf("Expected size: %d, got %d", testQueueSize, size)
	}
}

func getEmptyQueue(size int) *Queue {
	return New(size)
}

func getFillQueue(size int) *Queue {
	s := New(size)

	for i := 1; i <= testQueueSize; i++ {
		_ = s.Insert(i)
	}

	return s
}
