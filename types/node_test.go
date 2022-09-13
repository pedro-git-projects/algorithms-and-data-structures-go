package types

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	got := NewNode(1)
	want := Node{
		value: 1,
		next:  nil,
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("expected %v, but got %v", got, want)
	}
}
