package stack

type node[T any] struct {
	item T
	next *node[T]
}

type Stack[T any] interface {
	Push(T)
	Pop() T
	IsEmpty() bool
}

type stack[T any] struct {
	next *node[T]
}

func (s *stack[T]) Push(item T) {
	n := &node[T]{
		item: item,
	}
	n.next = s.next
	s.next = n
}

func (s *stack[T]) Pop() T {
	n := s.next
	s.next = s.next.next
	return n.item
}

func (s *stack[T]) IsEmpty() bool {
	return s.next == nil
}

func NewStack[T any]() Stack[T] {
	s := &stack[T]{
		next: nil,
	}

	return s
}
