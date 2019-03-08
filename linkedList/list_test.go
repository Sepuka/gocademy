package linkedList

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	list := linkedList{}
	if list.len != 0 {
		t.Error("empty list len not equal zero")
	}

	keyStack := []int{1,2}

	list.toHead(NewNode(keyStack[0]))
	if list.len != 1 {
		t.Error("new element did not add to the list")
	}

	list.toHead(NewNode(keyStack[1]))
	if list.len != 2 {
		t.Error("new element did not add to the list")
	}

	for i, node := range list.list() {
		pos := len(keyStack) - 1 - i
		if node.key != keyStack[pos] {
			t.Errorf("wrong order: expected %d, got %d", keyStack[pos], node.key)
		}
	}
}
