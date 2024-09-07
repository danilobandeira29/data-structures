package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	l := NewLinkedList()
	FIRST_ELEMENT := 44
	l.Push(FIRST_ELEMENT)
	l.Push(22)
	l.Push(33)
	if l.IsEmpty() {
		t.Error("linkedlist: expected to have elements\n")
	}
	e := l.Pop()
	if e != FIRST_ELEMENT {
		t.Errorf("linkedlist: first element removed expected %d got instead %d\n", FIRST_ELEMENT, e)
	}
	elements := l.List()
	expected := []int{22, 33}
	for idx := range elements {
		if elements[idx] != expected[idx] {
			t.Errorf("linkedlist: elements expected %v got instead %v\n", expected, elements)
		}
	}
}
