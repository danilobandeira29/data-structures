package hashmap

type Node struct {
	key   int
	value string
	next  *Node
}

func NewNode(key int, value string) *Node {
	return &Node{
		key:   key,
		value: value,
		next:  nil,
	}
}

func (n *Node) HashCode() int {
	return n.key % 100
}

func (n Node) Value() string {
	return n.value
}

// Hashmap
// this implementation create a Node when has collision, and get last Node in that Key to be the next last element
type HashMap struct {
	nodes []*Node
}

func NewHashMap() *HashMap {
	return &HashMap{
		nodes: make([]*Node, 200),
	}
}

func (h *HashMap) Put(n *Node) {
	register := h.nodes[n.HashCode()]
	if register == nil {
		h.nodes[n.HashCode()] = n
		return
	}
	if register.key == n.key {
		h.nodes[n.HashCode()].value = n.value
		return
	}
	for register.next != nil {
		register = register.next
	}
	h.nodes[n.HashCode()].next = n
}

func (h HashMap) Get(k int) *Node {
	n := NewNode(k, "")
	register := h.nodes[n.HashCode()]
	if register != nil && register.key == n.key {
		return register
	}
	for register.next != nil {
		if register.key == n.key {
			return register
		}
		register = register.next
	}
	return register
}

// func main() {
// 	n1, n2, n3, n4 := hashmap.NewNode(2245, "mensagem node 1"),
// 		hashmap.NewNode(2224, "mensagem node 2"),
// 		hashmap.NewNode(4567, "mensagem node 3"),
// 		hashmap.NewNode(200, "mensagem node 4")
// 	n5 := hashmap.NewNode(999, "nova mensagem pro node 1")
// 	n6 := hashmap.NewNode(99999, "same index")
// 	n7 := hashmap.NewNode(999, "novo novo, zerado")
// 	hashMap := hashmap.NewHashMap()
// 	hashMap.Put(n1)
// 	hashMap.Put(n2)
// 	hashMap.Put(n3)
// 	hashMap.Put(n4)
// 	hashMap.Put(n5)
// 	hashMap.Put(n6)
// 	hashMap.Put(n7)
// 	fmt.Println(hashMap.Get(999))
// }
