package tree

type Node struct {
	value int
}

func NewNode(v int) *Node {
	return &Node{
		value: v,
	}
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) SetValue(v int) {
	n.value = v
}
