package doublylinkedlist

type node struct {
	value    int
	next     *node
	previous *node
}

func newNode(v int) *node {
	return &node{
		value:    v,
		next:     nil,
		previous: nil,
	}
}
