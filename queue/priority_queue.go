// Priority Queue data structure with fixed size
// Algorithmic complexity:
// Insert O(N)
// Remove O(1)
package queue

import (
	"errors"
)

type PriorityQueue struct {
	numElem int
	queue   []int
}

func NewPriorityQueue(size int) *PriorityQueue {
	return &PriorityQueue{numElem: 0, queue: make([]int, size)}
}

func (q *PriorityQueue) Insert(v int) error {
	if q.isFull() {
		return errors.New("queue is full")
	}

	if q.isEmpty() {
		q.queue[q.numElem] = v
		q.numElem++

		return nil
	}

	var i int
	for i = q.numElem - 1; i >= 0; i-- {
		if v > q.queue[i] {
			q.queue[i+1] = q.queue[i]
		} else {
			break
		}
	}
	q.queue[i+1] = v
	q.numElem++

	return nil
}

func (q *PriorityQueue) Remove() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("queue is empty")
	}
	q.numElem--

	return q.queue[q.numElem], nil
}

func (q *PriorityQueue) Size() int {
	return q.numElem
}

func (q *PriorityQueue) isEmpty() bool {
	return q.numElem == 0
}

func (q *PriorityQueue) isFull() bool {
	return len(q.queue) == q.numElem
}
