package linkedlist

import (
	"errors"
)

type StackAsLinkedList struct {
	top *Node
}

var (
	ErrEmptyPop = errors.New("stackaslinkedlist: list must not be empty in order to pop")
)

func NewStackAsLinkedList() *StackAsLinkedList {
	return &StackAsLinkedList{
		top: nil,
	}
}

func (s *StackAsLinkedList) Push(n int) {
	node := NewNode(n)
	node.WriteNext(s.top)
	s.top = node
}

func (s *StackAsLinkedList) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, ErrEmptyPop
	}
	v := s.top.value
	s.top = s.top.Next()
	return v, nil
}

func (s *StackAsLinkedList) IsEmpty() bool {
	return s.top == nil
}
