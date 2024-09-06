package linkedlist

import (
	"errors"
)

type StackAsLinkedList struct {
	// do not need top, because start already solve the problem to run the linked list. branch for pull request refactored/stackaslinkedlist
	top   *Node
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
		top:   nil,
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
	}
	s.top = node
	s.total++
	return nil
}

func (s *StackAsLinkedList) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, ErrEmptyPop
	}
	v := s.top.value
	aux := s.start
	for aux.Next() != nil {
		aux = aux.Next()
	}
	s.top = aux
	s.total--
	return v, nil
}

func (s *StackAsLinkedList) IsFull() bool {
	return s.total == MAX_SIZE
}

func (s *StackAsLinkedList) IsEmpty() bool {
	return s.top == nil
}
