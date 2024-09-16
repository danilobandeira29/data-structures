package tree

import "fmt"

type Tree struct {
	element *Node
	left    *Tree
	right   *Tree
}

func NewTree(n *Node) *Tree {
	return &Tree{
		element: n,
		left:    nil,
		right:   nil,
	}
}

func (t *Tree) IsEmpty() bool {
	return t.element == nil
}

func (t *Tree) Insert(newNode *Node) {
	if t.IsEmpty() {
		t.element = newNode
		return
	}
	tree := NewTree(newNode)
	if tree.element.Value() < t.element.Value() && t.left == nil {
		t.left = tree
		fmt.Printf("insert element %d as left %d\n", newNode.Value(), t.element.Value())
		return
	}
	if tree.element.Value() < t.element.Value() && t.left != nil {
		t.left.Insert(newNode)
		return
	}
	if tree.element.Value() > t.element.Value() && t.right == nil {
		t.right = tree
		fmt.Printf("insert element %d as right %d\n", newNode.Value(), t.element.Value())
		return
	}
	if tree.element.Value() > t.element.Value() && t.right != nil {
		t.right.Insert(newNode)
		return
	}
}

func (t *Tree) Exists(v int) bool {
	if t.IsEmpty() {
		return false
	}
	if t.element.Value() == v {
		return true
	}
	if v < t.element.Value() && t.left == nil {
		return false
	}
	if v < t.element.Value() && t.left != nil {
		return t.left.Exists(v)
	}
	if v > t.element.Value() && t.right == nil {
		return false
	}
	if v > t.element.Value() && t.right != nil {
		return t.right.Exists(v)
	}
	return false
}
