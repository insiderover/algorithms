package linked_list

import "fmt"

type node struct {
	Key   int
	Value int
	next  *node
}

func NewNode(key int, value int) *node {
	return &node{Key: key, Value: value}
}

func (n *node) String() string {
	return fmt.Sprintf("|%d:%d|", n.Key, n.Value)
}
