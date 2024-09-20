package tree_test

import (
	"bytes"
	tree "github.com/danilobandeira29/data-structures/10_tree"
	"testing"
)

func TestNewTree(t *testing.T) {
	elements := []int{5, 1, 8, 15, 12, 18}
	tri := tree.NewTree(tree.NewNode(10))
	for _, e := range elements {
		tri.Insert(tree.NewNode(e))
	}
	for _, e := range elements {
		if !tri.Exists(e) {
			t.Errorf("tree: expect %d to exist\n", e)
			return
		}
	}
	if tri.Exists(13) {
		t.Errorf("tree: expected 13 to not exist")
		return
	}
}

func TestTree_InOrder(t *testing.T) {
	var buf bytes.Buffer
	tri := tree.NewTree(tree.NewNode(10))
	elements := []int{5, 1, 8, 15, 12, 18}
	for _, e := range elements {
		tri.Insert(tree.NewNode(e))
	}
	tri.InOrder(&buf)
	if buf.String() != "1 5 8 10 12 15 18 " {
		t.Errorf("not in order, received %s\n", buf.String())
	}
}
