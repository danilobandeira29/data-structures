package stack

// Stack
// Last In First Out - LIFO
type Stack[T any] struct {
	values [100]T
	top    int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		values: [100]T{},
		top:    -1,
	}
}

func (s *Stack[T]) Push(e T) {
	s.top++
	s.values[s.top] = e
}

func (s *Stack[T]) Pop() T {
	e := s.values[s.top]
	s.top--
	return e
}

func (s *Stack[T]) IsFull() bool {
	return s.top == 9
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == -1
}
