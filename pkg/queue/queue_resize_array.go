package queue

import "sync"

type queueRA[T any] struct {
	sync.Mutex
	first int
	last  int
	array []T
}

func (q *queueRA[T]) Enqueue(item T) {
	q.Lock()
	defer q.Unlock()
	q.last++
	if q.last >= len(q.array) {
		// Resize, double the length
		q.array = resize(q.array, q.last*2)
	}
	q.array[q.last-1] = item
}

func (q *queueRA[T]) Dequeue() T {
	q.Lock()
	defer q.Unlock()
	if q.last-q.first <= len(q.array)/4 {
		q.array = resize(q.array, (q.last-q.first)/2)
	}
	q.first++
	return q.array[q.first-1]
}

func (q *queueRA[T]) IsEmpty() bool {
	q.Lock()
	defer q.Unlock()
	return q.first == q.last
}

func NewQueueRA[T any]() Queue[T] {
	q := &queueRA[T]{
		array: make([]T, 0),
	}

	return q
}

func resize[T any](s []T, l int) []T {
	if l <= len(s) {
		return s
	}
	ns := make([]T, l)
	copy(ns, s)

	return ns
}
