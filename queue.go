package main

type queue chan Square

func (q queue) Push(s Square){
	q <- s
}

func (q queue) Pop() Square{
	return <- q
}

func (q queue) Length() int{
	return len(q)
}