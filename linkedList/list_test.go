package linkedList

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	expectedList := []int{2,1,3,4}

	list := linkedList{}
	if list.len != 0 {
		t.Error("empty list len not equal zero")
	}

	list.toHead(NewNode(1))
	list.toTail(NewNode(3))
	list.toHead(NewNode(2))
	list.toTail(NewNode(4))

	if list.len != len(expectedList) {
		t.Error("wrong list length")
	}

	for i, node := range list.list() {
		if node.key != expectedList[i] {
			t.Errorf("wrong order: expected %d, got %d", expectedList[i], node.key)
		}
	}
}
