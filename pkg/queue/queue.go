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

func (s *queue[T]) Enqueue(item T) {
	n := &node[T]{
		item: item,
	}
	s.Lock()
	defer s.Unlock()
	if s.first == nil {
		s.first = n
		s.last = s.first
	} else {
		s.last.next = n
		s.last = n
	}
}

func (s *queue[T]) Dequeue() T {
	n := s.first
	if s.first == nil {
		return *new(T)
	}
	s.Lock()
	defer s.Unlock()
	s.first = s.first.next

	return n.item
}

func (s *queue[T]) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()

	return s.first == nil
}

func NewQueue[T any]() Queue[T] {
	s := &queue[T]{
		first: nil,
	}

	return s
}
