package hashmap_test

import (
	"testing"

	hashmap "github.com/danilobandeira29/data-structures/09_hash_map"
)

func TestHashMap(t *testing.T) {
	hashMap := hashmap.NewHashMap()
	n1 := hashmap.NewNode(1, "first value")
	hashMap.Put(n1)
	result := hashMap.Get(1)
	if result.Value() != "first value" {
		t.Errorf("hashmap: expect mgs to be \"first value\" but got %s", result.Value())
	}
	oldValue := result.Value()
	newN1 := hashmap.NewNode(1, "new value for n1")
	hashMap.Put(newN1)
	result = hashMap.Get(1)
	if result.Value() == oldValue {
		t.Error("hashmap: value of node not updated")
	}
	n2 := hashmap.NewNode(99999, "node 99999")
	nodeWithSameHashCode := hashmap.NewNode(999, "node 999")
	hashMap.Put(n2)
	hashMap.Put(nodeWithSameHashCode)
	result = hashMap.Get(99999)
	if result.Value() != "node 99999" {
		t.Error("hashmap: the node with same key replaced, but should not do that")
	}
	result = hashMap.Get(999)
	if result.Value() != "node 999" {
		t.Error("hashmap: unreferenced node")
	}
}
