package types

type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	n := new(Node)
	n.value = value
	n.next = nil
	return n
}
