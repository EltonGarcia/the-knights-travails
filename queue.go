package main

type queue chan Path

func (q queue) Push(s Path) {
	q <- s
}

func (q queue) Pop() Path {
	return <-q
}

func (q queue) Length() int {
	return len(q)
}

func (q queue) IsEmpty() bool {
	return q.Length() == 0
}
