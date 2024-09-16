package tree_test

import (
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
