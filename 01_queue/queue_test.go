package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	if !q.IsEmpty() {
		t.Error("queue must init empty")
	}
	elements := []int{10, 5, 11, 55}
	for !q.IsFull() {
		for _, e := range elements {
			q.Enqueue(e)
		}
	}
	if !q.IsFull() {
		t.Error("queue must be full")
	}
	for !q.IsEmpty() {
		for _, e := range elements {
			removed := q.Dequeue()
			if removed != e {
				t.Errorf("element expected to be %d but return %d", e, removed)
			}
		}
	}
	if !q.IsEmpty() {
		t.Error("all elements removed, expected queue to be empty")
	}
	if q.IsFull() {
		t.Error("queue cannot be full")
	}
}
