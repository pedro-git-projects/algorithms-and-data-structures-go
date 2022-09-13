package types

import (
	"reflect"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	got := NewLinkedList(1)
	n := NewNode(1)

	want := LinkedList{
		head:   n,
		tail:   n,
		length: 1,
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("expected %v but got %v", got, want)
	}
}

func TestAppend(t *testing.T) {
	// nil list
	nll := new(LinkedList)
	nll.Append(1)
	got := nll
	want := NewLinkedList(1)
	if !reflect.DeepEqual(*got, *want) {
		t.Errorf("expected %v but got %v", got, want)
	}

	// populated list
	/*
		ll := NewLinkedList(1)
		ll.Append(2)
		n1 := NewNode(1)
		n2 := NewNode(2)
		want2 := LinkedList{
			head:   n1,
			tail:   n2,
			length: 2,
		}
		if *got != want2 {
			t.Errorf("expected %v but got %v", got, want)
		}
	*/
}
