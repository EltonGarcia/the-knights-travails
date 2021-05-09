package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	path1 := newPath(Square{X: 1, Y: 1})
	path2 := newPath(Square{X: 2, Y: 2})

	queue := make(Queue, 2)

	assert.Equal(t, true, queue.IsEmpty())

	queue.Enqueue(path1)
	queue.Enqueue(path2)

	assert.Equal(t, 2, queue.Size())
	assert.Equal(t, false, queue.IsEmpty())

	assert.Equal(t, path1, queue.Dequeue())
	assert.Equal(t, path2, queue.Dequeue())

	assert.Equal(t, true, queue.IsEmpty())
}
