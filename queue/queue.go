// Queue data structure with fixed size
// Algorithmic complexity:
// Insert O(1)
// Remove O(1)
package queue

import (
	"errors"
)

type Queue struct {
	rear    int
	front   int
	numElem int
	queue   []int
}

func NewQueue(size int) *Queue {
	return &Queue{rear: -1, front: 0, numElem: 0, queue: make([]int, size)}
}

func (q *Queue) Insert(v int) error {
	if q.isFull() {
		return errors.New("queue is full")
	}

	if q.rear == len(q.queue)-1 {
		q.rear = -1
	}

	q.rear++
	q.queue[q.rear] = v
	q.numElem++

	return nil
}

func (q *Queue) Remove() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("queue is empty")
	}

	tmp := q.queue[q.front]
	q.front++
	if q.front == len(q.queue) {
		q.front = 0
	}
	q.numElem -= 1

	return tmp, nil
}

func (q *Queue) Size() int {
	return q.numElem
}

func (q *Queue) isEmpty() bool {
	return q.numElem == 0
}

func (q *Queue) isFull() bool {
	return len(q.queue) == q.numElem
}
