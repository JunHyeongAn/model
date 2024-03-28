package model

import "container/list"

type Queue struct {
	l *list.List
}

func New() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Push(v interface{}) {
	q.l.PushBack(v)
}

func (q *Queue) Pop() {
	front := q.l.Front()

	q.l.Remove(front)
}
