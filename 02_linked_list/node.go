package linkedlist

type Node struct {
	value int
	next  *Node
}

func NewNode(e int) *Node {
	return &Node{
		value: e,
		next:  nil,
	}
}

func (n *Node) WriteNext(newNode *Node) {
	n.next = newNode
}

func (n *Node) Next() *Node {
	return n.next
}
