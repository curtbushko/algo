package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	// Create a new stack
	q := NewQueue[int]()

	// Verify that the stack is empty
	assert.True(t, q.IsEmpty(), "expected queue to be empty")

	// Verify that the length of the queue is 0
	assert.Equal(t, 0, q.Len(), "expected queue length to be 0")

	q.Push(1)
	q.Push(2)
	q.Push(3)

	// Verify that the queue is not empty
	assert.False(t, q.IsEmpty(), "expected queue to be not empty")

	// Verify that the length of the queue is 3
	assert.Equal(t, 3, q.Len(), "expected queue length to be 3")

	// Verify that Peek() returns the first item in the queue
	assert.Equal(t, 1, q.Peek(), "expected Peek() to return 1")

	// Verify that PushN() adds multiple item on the queue
	q.PushN([]int{4, 5, 6})
	assert.Equal(t, 6, q.Len(), "expected queue length to be 6 after PushN()")

	// Verify that Pop() removes and returns the top item of the queue
	item := q.Pop()
	assert.Equal(t, 1, item, "expected Pop() to return 1")
	assert.Equal(t, 5, q.Len(), "expected queue length to be 5 after Pop()")

	// Verify that PopN() removes and returns multiple items from the queue
	items, err := q.PopN(3)
	fmt.Println(items)
	assert.Nil(t, err, "expected PopN() to return itemsents from the stack")
	assert.Len(t, items, 3, "expected PopN() to return 3 itemsents")
	assert.Equal(t, []int{2, 3, 4}, items, "expected Pop() to return 1")
	assert.Equal(t, 3, q.Len(), "expected stack length to be 2 after PopN()")
}
