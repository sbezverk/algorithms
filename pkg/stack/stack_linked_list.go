package stack

import "sync"

type node[T any] struct {
	item T
	next *node[T]
}

type StackLL[T any] interface {
	Push(T)
	Pop() T
	IsEmpty() bool
}

type stackLL[T any] struct {
	sync.Mutex
	next *node[T]
}

func (s *stackLL[T]) Push(item T) {
	n := &node[T]{
		item: item,
	}
	s.Lock()
	defer s.Unlock()
	n.next = s.next
	s.next = n
}

func (s *stackLL[T]) Pop() T {
	n := s.next
	if s.next == nil {
		return *new(T)
	}
	s.Lock()
	defer s.Unlock()
	s.next = s.next.next

	return n.item
}

func (s *stackLL[T]) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()

	return s.next == nil
}

func NewStackLL[T any]() StackLL[T] {
	s := &stackLL[T]{
		next: nil,
	}

	return s
}
