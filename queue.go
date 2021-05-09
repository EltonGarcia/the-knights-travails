package main

// Queue implementation using a channel
type Queue chan Path

// Enqueue adds an item to the queue.
func (q *Queue) Enqueue(s Path) {
	*q <- s
}

// Dequeue removes and returns the
// item in the beginning of the Queue.
func (q *Queue) Dequeue() Path {
	return <-*q
}

// Size counts the items in the Queue.
func (q *Queue) Size() int {
	return len(*q)
}

// IsEmpty checks if a Queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
