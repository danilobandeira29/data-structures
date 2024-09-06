package linkedlist

import (
	"errors"
)

type StackAsLinkedList struct {
	start *Node
	total int
}

var (
	MAX_SIZE     = 10
	ErrStackFull = errors.New("stackaslinkedlist: stack is full")
	ErrEmptyPop  = errors.New("stackaslinkedlist: list must not be empty in order to pop")
)

func NewStackAsLinkedList() *StackAsLinkedList {
	return &StackAsLinkedList{
		start: nil,
		total: 0,
	}
}

func (s *StackAsLinkedList) Push(n int) error {
	if s.IsFull() {
		return ErrStackFull
	}
	node := NewNode(n)
	if s.IsEmpty() {
		s.start = node
		return nil
	}
	aux := s.start
	for aux.Next() != nil {
		aux = aux.Next()
	}
	aux.WriteNext(node)
	s.total++
	return nil
}

func (s *StackAsLinkedList) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, ErrEmptyPop
	}
	aux := s.start
	for aux.Next() != nil {
		aux = aux.Next()
	}
	v := aux.value
	s.total--
	return v, nil
}

func (s *StackAsLinkedList) IsFull() bool {
	return s.total == MAX_SIZE
}

func (s *StackAsLinkedList) IsEmpty() bool {
	return s.start == nil
}
