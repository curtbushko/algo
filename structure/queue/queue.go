package queue

import "fmt"

// Queue is a generic queue that is FIFO or LILO
//
//	Usage:
//	q := Queue([]int{1,2,3})
//	q := NewQueue[int]()
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new empty queue
// Usage:
// s := Stack([]int{1,2,3})
// s := NewStack[int]()
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: []T{},
	}
}

// IsEmpty checks to see if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Len returns the length of the queue
func (q *Queue[T]) Len() int {
	return len(q.items)
}

// PushN pushes a number of items onto the queue
func (q *Queue[T]) PushN(v []T) {
	q.items = append(q.items, v...)
}

// Push adds an item to the top of the queue
func (q *Queue[T]) Push(v T) {
	q.items = append(q.items, v)
}

// PopN pops the number of items off of the queue. Return an error if popping more items than available.
func (q *Queue[T]) PopN(n int) ([]T, error) {
	var r []T
	if q.Len() < n {
		return r, fmt.Errorf("cannot pop %d items from stack of length %d", n, q.Len())
	}

	r = q.items[q.Len()-n:]
	q.items = q.items[q.Len()-n:] // subject to memory leak
	return r, nil
}

// Pop pops a single item off of front of the queue
func (q *Queue[T]) Pop() any {
	if q.IsEmpty() {
		return nil
	}
	i := q.items[0]
	q.items = q.items[1:]
	return i
}

func (q *Queue[T]) Peek() any {
	if q.IsEmpty() {
		return nil
	}
	return q.items[0]
}
