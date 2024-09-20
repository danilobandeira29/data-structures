package tree

import (
	"fmt"
	"io"
)

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

// Remove
// returns a new tree if find node, if not return same tree
func (t *Tree) Remove(n *Node) *Tree {
	if t.element.Value() == n.Value() {
		// case simple
		if t.left == nil && t.right == nil {
			return nil
		}
		// case 2
		if t.left != nil && t.right == nil {
			return t.left
		}
		// case 3
		if t.right != nil && t.left == nil {
			return t.right
		}
		// worst case
		if t.left != nil && t.right != nil {
			// get greater element from left sub-tree
			tri := t.left
			for tri.right != nil {
				tri = tri.right
			}
			t.element, tri.element = tri.element, n
			t.left = t.left.Remove(n)
		}
	}
	if n.Value() < t.element.Value() && t.left != nil {
		t.left = t.left.Remove(n)
	}
	if n.Value() > t.element.Value() && t.right != nil {
		t.right = t.right.Remove(n)
	}
	return t
}

func (t *Tree) PreOrder() {
	if t.IsEmpty() {
		return
	}
	fmt.Print(t.element.Value(), " ")
	if t.left != nil {
		t.left.PreOrder()
	}
	if t.right != nil {
		t.right.PreOrder()
	}
}

func (t *Tree) InOrder(r io.Writer) {
	if t.IsEmpty() {
		return
	}
	if t.left != nil {
		t.left.InOrder(r)
	}
	v := fmt.Sprintf("%d ", t.element.Value())
	_, err := r.Write([]byte(v))
	if err != nil {
		panic(err)
	}
	if t.right != nil {
		t.right.InOrder(r)
	}
}

func (t *Tree) PosOrder() {
	if t.IsEmpty() {
		return
	}
	if t.right != nil {
		t.right.PosOrder()
	}
	if t.left != nil {
		t.left.PosOrder()
	}
	fmt.Print(t.element.Value(), " ")
}

func (t *Tree) InOrderReverse() {
	if t.IsEmpty() {
		return
	}
	if t.right != nil {
		t.right.InOrderReverse()
	}
	fmt.Print(t.element.Value(), " ")
	if t.left != nil {
		t.left.InOrderReverse()
	}
}
