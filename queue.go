package cache

import (
	"container/list"
	"sync"
)

type queue struct {
	*list.List
	lock sync.Mutex
}

func NewQueue() *queue {
	return &queue{
		List: list.New(),
		lock: sync.Mutex{},
	}
}

func (q *queue) PushFront(v interface{}) {
	defer q.lock.Unlock()
	q.lock.Lock()
	q.List.PushFront(v)
}

func (q *queue) Remove(e *list.Element) {
	defer q.lock.Unlock()
	q.lock.Lock()
	q.List.Remove(e)
}

func (q *queue) MoveToFront(e *list.Element) {
	defer q.lock.Unlock()
	q.lock.Lock()
	q.List.MoveToFront(e)
}
