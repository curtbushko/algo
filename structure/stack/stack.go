package stack

import "fmt"

// Stack is a generic stack of that is LIFO or FILO
//
//	Usage:
//	s := Stack([]int{1,2,3})
//	s := NewStack[int]()
type Stack[T any] struct {
	items []T
}

// NewStack creates a new empty stack
// Usage:
// s := Stack([]int{1,2,3})
// s := NewStack[int]()
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: []T{},
	}
}

// IsEmpty checks to see if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Len returns the length of the stack
func (s *Stack[T]) Len() int {
	return len(s.items)
}

// Peek looks at the item at the top of the stack. Returns nil if empty.
func (s *Stack[T]) Peek() any {
	if s.IsEmpty() {
		return nil
	}
	return s.items[s.Len()-1]
}

// PushN pushes a number of items onto the stack
func (s *Stack[T]) PushN(v []T) {
	s.items = append(s.items, v...)
}

// Push pushes a single item on the stack
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// PopN pops the number of items off of the stack. Return an error if popping more items than available.
func (s *Stack[T]) PopN(n int) ([]T, error) {
	var r []T
	if s.Len() < n {
		return r, fmt.Errorf("cannot pop %d items from stack of length %d", n, s.Len())
	}

	r = s.items[s.Len()-n:]
	s.items = s.items[:s.Len()-n] // subject to memory leak
	return r, nil
}

// Pop pops a single item off of the stack
func (s *Stack[T]) Pop() any {
	if s.IsEmpty() {
		return nil
	}
	l := s.Len() - 1
	i := s.items[l]
	s.items = s.items[:l]
	return i
}
