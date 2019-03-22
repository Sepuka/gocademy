package linkedList

import (
	"testing"
)

func TestToTail(t *testing.T) {
	list := linkedList{}
	list.toTail(NewNode(1))
	list.toTail(NewNode(2))
	list.toTail(NewNode(3))
	expected := map[int]int{
		0: 1,
		1: 2,
		2: 3,
	}
	for i, v := range list.list() {
		if expected[i] != v.key {
			t.Errorf("Expected %d got %d", expected[i], v.key)
		}
	}
}

func TestToHead(t *testing.T) {
	list := linkedList{}
	list.toHead(NewNode(1))
	list.toHead(NewNode(2))
	list.toHead(NewNode(3))
	expected := map[int]int{
		0: 3,
		1: 2,
		2: 1,
	}
	for i, v := range list.list() {
		if expected[i] != v.key {
			t.Errorf("Expected %d got %d", expected[i], v.key)
		}
	}
}

func TestMixedInserts(t *testing.T) {
	expectedList := []int{2,1,5,3,4}

	list := linkedList{}
	if list.len != 0 {
		t.Error("empty list len not equal zero")
	}

	list.toHead(NewNode(1))
	list.toTail(NewNode(3))
	list.toHead(NewNode(2))
	list.toTail(NewNode(4))
	list.toPos(NewNode(5), 2)

	if list.len != len(expectedList) {
		t.Error("wrong list length")
	}

	for i, node := range list.list() {
		if node.key != expectedList[i] {
			t.Errorf("wrong order: expected %d, got %d", expectedList[i], node.key)
		}
	}
}
