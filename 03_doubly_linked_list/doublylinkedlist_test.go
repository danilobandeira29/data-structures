package doublylinkedlist

import (
	"slices"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.Insert(44)
	dll.Insert(55)
	dll.Insert(33)
	dll.Insert(22)
	if !slices.Equal(dll.InOrder(), []int{44, 55, 33, 22}) {
		t.Error("doublylinkedlist: not in order")
	}
	if !slices.Equal(dll.Reverse(), []int{22, 33, 55, 44}) {
		t.Error("doublylinkedlist: not in reverse order")
	}
	element, err := dll.RemoveFromStart()
	if err != nil {
		t.Error("doublylinkedlist: not possible to remove from start")
	}
	if element != 44 {
		t.Errorf("doublylinkedlist: expect %d but got %d\n", 44, element)
	}
	element, err = dll.RemoveFromEnd()
	if err != nil {
		t.Error("doublylinkedlist: not possible to remove from end")
	}
	if element != 22 {
		t.Errorf("doublylinkedlist: expect %d but got %d\n", 22, element)
	}
}
