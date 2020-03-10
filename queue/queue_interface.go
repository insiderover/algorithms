package queue

type IQueue interface {
	Insert(v int) error
	Remove() (int, error)
	Size() int
}
