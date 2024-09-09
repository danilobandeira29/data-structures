package doublylinkedlist

import (
	"errors"
)

type DoublyLinkedList struct {
	startNode *node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		startNode: nil,
	}
}

func (dll DoublyLinkedList) IsEmpty() bool {
	return dll.startNode == nil
}

func (dll DoublyLinkedList) InOrder() []int {
	aux := dll.startNode
	result := []int{}
	for aux.next != dll.startNode {
		result = append(result, aux.value)
		aux = aux.next
	}
	result = append(result, aux.value)
	return result
}

func (dll DoublyLinkedList) Reverse() []int {
	aux := dll.startNode
	result := []int{}
	for aux.previous != dll.startNode {
		result = append(result, aux.previous.value)
		aux = aux.previous
	}
	result = append(result, aux.previous.value)
	return result
}

func (dll *DoublyLinkedList) RemoveFromStart() (int, error) {
	if dll.IsEmpty() {
		return -1, errors.New("doublylinkedlist: empty")
	}
	value := dll.startNode.value
	if dll.startNode.next == nil {
		return value, nil
	}
	aux := dll.startNode
	for aux.next != dll.startNode {
		aux = aux.next
	}
	dll.startNode = dll.startNode.next
	dll.startNode.previous = aux
	aux.next = dll.startNode
	return value, nil
}

func (dll *DoublyLinkedList) RemoveFromEnd() (int, error) {
	if dll.IsEmpty() {
		return -1, errors.New("doublylinkedlist: empty")
	}
	aux := dll.startNode
	for aux.next != dll.startNode {
		aux = aux.next
	}
	dll.startNode.previous = aux.previous
	aux.previous.next = dll.startNode
	return aux.value, nil
}

func (dll *DoublyLinkedList) Insert(v int) {
	node := newNode(v)
	if dll.IsEmpty() {
		dll.startNode = node
		node.next = node
		node.previous = node
		return
	}
	aux := dll.startNode
	for dll.startNode != aux.next {
		aux = aux.next
	}
	aux.next = node
	dll.startNode.previous = node
	node.next = dll.startNode
	node.previous = aux
}
