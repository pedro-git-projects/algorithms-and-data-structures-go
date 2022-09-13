package types

import "fmt"

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList(value int) *LinkedList {
	n := NewNode(value)
	ll := new(LinkedList)
	ll.head = n
	ll.tail = n
	ll.length = 1
	return ll
}

func (ll LinkedList) PrintLinkedList() {
	temp := ll.head
	for temp != nil {
		fmt.Printf("%d", temp.value)
		temp = temp.next
	}
	fmt.Printf("\n")
}

func (ll *LinkedList) Append(value int) {
	n := NewNode(value)
	if ll.head != nil {
		ll.tail.next = n
		ll.tail = n
	} else {
		ll.head = n
		ll.tail = n
	}
	ll.length++
}

func (ll *LinkedList) Prepend(value int) {
	n := NewNode(value)
	if ll.head != nil {
		n.next = ll.head
		ll.head = n
	} else {
		ll.head = n
		ll.tail = n
		n.next = nil
	}
	ll.length++
}
