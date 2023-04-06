package stack

import "sync"

type stackRA[T any] struct {
	sync.Mutex
	n     int
	stack []T
}

func (s *stackRA[T]) Push(item T) {
	s.Lock()
	defer s.Unlock()
	s.n++
	if s.n >= len(s.stack) {
		// Resize, double the length
		s.stack = resize(s.stack, s.n*2)
	}
	s.stack[s.n-1] = item
}

func (s *stackRA[T]) Pop() T {
	if s.n <= 0 {
		return *new(T)
	}
	if s.n == len(s.stack)/4 {
		// Resize, cut the stack size in 2
		s.stack = resize(s.stack, len(s.stack)/2)
	}
	s.n--

	return s.stack[s.n]
}

func (s *stackRA[T]) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()
	return s.n == 0
}

func NewStackRA[T any]() Stack[T] {
	s := &stackRA[T]{
		stack: make([]T, 0),
	}

	return s
}

func resize[T any](s []T, l int) []T {
	if l <= len(s) {
		return s
	}
	ns := make([]T, l)
	copy(ns, s)

	return ns
}
