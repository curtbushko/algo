package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	// Create a new stack
	s := NewStack[int]()

	// Verify that the stack is empty
	assert.True(t, s.IsEmpty(), "expected stack to be empty")

	// Verify that the length of the stack is 0
	assert.Equal(t, 0, s.Len(), "expected stack length to be 0")

	// Push some itemsents onto the stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Verify that the stack is not empty
	assert.False(t, s.IsEmpty(), "expected stack to be not empty")

	// Verify that the length of the stack is 3
	assert.Equal(t, 3, s.Len(), "expected stack length to be 3")

	// Verify that Peek() returns the top item of the stack
	assert.Equal(t, 3, s.Peek(), "expected Peek() to return 3")

	// Verify that PushN() adds multiple items to the stack
	s.PushN([]int{4, 5, 6})
	assert.Equal(t, 6, s.Len(), "expected stack length to be 6 after PushN()")

	// Verify that Pop() removes and returns the top item of the stack
	item := s.Pop()
	assert.Equal(t, 6, item, "expected Pop() to return 6")
	assert.Equal(t, 5, s.Len(), "expected stack length to be 5 after Pop()")

	// Verify that PopN() removes and returns multiple items from the stack
	items, err := s.PopN(3)
	assert.Nil(t, err, "expected PopN() to return itemsents from the stack")
	assert.Len(t, items, 3, "expected PopN() to return 3 itemsents")
	assert.Equal(t, 2, s.Len(), "expected stack length to be 2 after PopN()")

	// Verify that PopN() returns an error if called with n > s.Len()
	_, err = s.PopN(3)
	assert.NotNil(t, err, "expected PopN() to return an error when called with n > s.Len()")
}
