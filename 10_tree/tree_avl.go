package tree

import (
	"fmt"
	"io"
)

type AVL struct {
	element *Node
	left    *AVL
	right   *AVL
	balance int
}

func NewTreeAVL() *AVL {
	return &AVL{
		element: nil,
		left:    nil,
		right:   nil,
		balance: 0,
	}
}

func (t *AVL) IsEmpty() bool {
	return t.element == nil
}

func (t *AVL) Insert(newNode *Node) *AVL {
	if t.IsEmpty() {
		t.element = newNode
		return t
	}
	tree := NewTreeAVL()
	tree.Insert(newNode)
	if tree.element.Value() < t.element.Value() && t.left == nil {
		t.left = tree
		fmt.Printf("insert element %d as left %d\n", newNode.Value(), t.element.Value())
	}
	if tree.element.Value() < t.element.Value() && t.left != nil {
		t.left = t.left.Insert(newNode)
	}
	if tree.element.Value() > t.element.Value() && t.right == nil {
		t.right = tree
		fmt.Printf("insert element %d as right %d\n", newNode.Value(), t.element.Value())
	}
	if tree.element.Value() > t.element.Value() && t.right != nil {
		t.right = t.right.Insert(newNode)
	}
	return t
}

func (t *AVL) Exists(v int) bool {
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
func (t *AVL) Remove(n *Node) *AVL {
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

func (t *AVL) PreOrder() {
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

func (t *AVL) InOrder(r io.Writer) {
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

func (t *AVL) PosOrder() {
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

func (t *AVL) InOrderReverse() {
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

func (t *AVL) Height() int {
	if t.left == nil && t.right == nil {
		return 1
	}
	if t.left != nil && t.right == nil {
		return 1 + t.left.Height()
	}
	if t.right != nil && t.left == nil {
		return 1 + t.right.Height()
	}
	return 1 + max(t.left.Height(), t.right.Height())
}

func (t *AVL) Balance() {
	if t.left == nil && t.right == nil {
		t.balance = 0
	}
	// right - left
	if t.left == nil && t.right != nil {
		t.balance = t.right.Height() - 0
	}
	if t.left != nil && t.right == nil {
		t.balance = 0 - t.left.Height()
	}
	if t.left != nil && t.right != nil {
		t.balance = t.right.Height() - t.left.Height()
	}
	if t.left != nil {
		t.left.Balance()
	}
	if t.right != nil {
		t.right.Balance()
	}
}

func (t *AVL) MakeBalance() *AVL {
	if t.balance >= 2 {
		if t.balance*t.right.balance > 0 {
			fmt.Println("simple rotate right")
			return t.SimpleRotateRight()
		}
		fmt.Println("double rotate right")
		return t.DoubleRotateRight()
	}
	if t.balance <= -2 {
		if t.balance*t.left.balance > 0 {
			fmt.Println("simple rotate left")
			return t.SimpleRotateLeft()
		}
		fmt.Println("double rotate left")
		return t.DoubleRotateLeft()
	}
	t.Balance()
	if t.left != nil {
		t.left = t.left.MakeBalance()
	}
	if t.right != nil {
		t.right = t.right.MakeBalance()
	}
	return t
}

func (t *AVL) SimpleRotateRight() *AVL {
	var tRight = t.right
	var tRightSon *AVL
	if t.right != nil {
		if t.right.left != nil {
			tRightSon = tRight.left
		}
	}
	tRight.left = t
	t.right = tRightSon
	return tRight
}

func (t *AVL) DoubleRotateRight() *AVL {
	var tri = t
	var tRight = t.right
	var tRightLeftSon = tRight.left
	var tInserted = tRightLeftSon.right
	tRight.left = tInserted
	tRightLeftSon.right = tRight
	t.right = tRightLeftSon
	var root = tri.right
	tri.right = nil
	root.left = tri
	return root
}

func (t *AVL) SimpleRotateLeft() *AVL {
	var tLeft = t.left
	var tLeftSon *AVL
	if t.left != nil {
		if t.left.right != nil {
			tLeftSon = tLeft.right
		}
	}
	t.left = tLeftSon
	tLeft.right = t
	return tLeft
}

func (t *AVL) DoubleRotateLeft() *AVL {
	var tri = t
	var tLeft = t.left
	var tLeftRightSon = tLeft.right
	var tInserted = tLeftRightSon.left
	tLeft.right = tInserted
	t.left = tLeftRightSon
	tLeftRightSon.left = tLeft
	var root = tri.left
	root.right = tri
	tri.left = nil
	return root
}

func (t *AVL) PrintBalance() {
	fmt.Printf("node %d has balance %d\n", t.element.Value(), t.balance)
}
