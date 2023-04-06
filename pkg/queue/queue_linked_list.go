package queue

import "sync"

type node[T any] struct {
	item T
	next *node[T]
}

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() T
	IsEmpty() bool
}

type queue[T any] struct {
	sync.Mutex
	first *node[T]
	last  *node[T]
}

func (q *queue[T]) Enqueue(item T) {
	n := &node[T]{
		item: item,
	}
	q.Lock()
	defer q.Unlock()
	if q.first == nil {
		q.first = n
		q.last = q.first
	} else {
		q.last.next = n
		q.last = n
	}
}

func (q *queue[T]) Dequeue() T {
	n := q.first
	if q.first == nil {
		return *new(T)
	}
	q.Lock()
	defer q.Unlock()
	q.first = q.first.next

	return n.item
}

func (q *queue[T]) IsEmpty() bool {
	q.Lock()
	defer q.Unlock()

	return q.first == nil
}

func NewQueue[T any]() Queue[T] {
	q := &queue[T]{
		first: nil,
	}

	return q
}
