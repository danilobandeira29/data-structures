package linkedlist

type LinkedList struct {
	firstElement *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		firstElement: nil,
	}
}

func (l *LinkedList) Insert(n int) {
	node := NewNode(n)
	if l.firstElement == nil {
		l.firstElement = node
		return
	}
	aux := l.firstElement
	for aux.Next() != nil {
		aux = aux.Next()
	}
	aux.WriteNext(node)
}

// Del expects that the client check if List is empty
func (l *LinkedList) Del() int {
	v := l.firstElement
	l.firstElement = l.firstElement.Next()
	return v.value
}

func (l LinkedList) List() []byte {
	result := []byte{}
	if l.firstElement == nil {
		return result
	}
	aux := l.firstElement
	for aux.Next() != nil {
		result = append(result, byte(aux.value))
		aux = aux.Next()
	}
	result = append(result, byte(aux.value))
	return result
}

func (l LinkedList) IsEmpty() bool {
	return l.firstElement == nil
}
