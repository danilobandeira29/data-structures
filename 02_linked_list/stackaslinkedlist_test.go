package linkedlist

import "testing"

func TestStackAsLinkedList(t *testing.T) {
	s := NewStackAsLinkedList()
	if !s.IsEmpty() {
		t.Error("stackaslinkedlist: stack must be empty when instantiate")
	}
	elements := []int{10, 22, 33}
	for _, e := range elements {
		s.Push(e)
	}
	v, err := s.Pop()
	if err == ErrEmptyPop {
		t.Error(err)
	}
	if v != 33 {
		t.Errorf("stackaslinkedlist: first element to remove expected to %d got instead %d\n", 33, v)
	}
	if s.IsEmpty() {
		t.Error("stackaslinkedlist: must have element, because not all elements has been poped")
	}
}
