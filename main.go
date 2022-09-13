package main

import (
	"data-structures/types"
)

func main() {
	ll := types.NewLinkedList(1)
	ll.Append(2)
	ll.Append(3)
	ll.PrintLinkedList()
	ll.Prepend(-1)
	ll.PrintLinkedList()
}
