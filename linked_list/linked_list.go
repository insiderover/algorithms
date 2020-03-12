package linked_list

import (
	"fmt"
)

type LinkedList struct {
	first *node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) InsertFirst(node *node) {
	if l.first == nil {
		l.first = node
		return
	}

	node.next = l.first
	l.first = node
}

func (l *LinkedList) DeleteFirst() *node {
	if l.IsEmpty() {
		return nil
	}

	tmp := l.first
	l.first = l.first.next

	return tmp
}

func (l *LinkedList) Find(key int) *node {
	if l.IsEmpty() {
		return nil
	}

	var current = l.first
	for current.Key != key {
		if current.next == nil {
			return nil
		}
		current = current.next
	}

	return current
}

func (l *LinkedList) Delete(key int) *node {
	var curr = l.first
	var prev = l.first

	for curr.Key != key {
		if curr.next == nil {
			return nil
		}

		prev = curr
		curr = curr.next
	}

	if l.first == curr {
		l.first = l.first.next
	}
	prev.next = curr.next

	return curr
}

func (l *LinkedList) IsEmpty() bool {
	return l.first == nil
}

func (l *LinkedList) String() string {
	var str string
	current := l.first
	for current != nil {
		str += fmt.Sprint(current) + " -> "

		current = current.next
	}

	str += fmt.Sprint(nil)

	return str
}
