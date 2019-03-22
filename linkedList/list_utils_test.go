package linkedList

import (
	"testing"
)

func TestNaiveReverse(t *testing.T) {
	list := linkedList{}
	for i := 0; i < 10; i++ {
		list.toTail(NewNode(i))
	}

	if list.tail != naiveInverse(list.head) {
		t.Errorf("The head of list must be equal the tail")
	}
}


func TestRecursiveReverse(t *testing.T) {
	list := linkedList{}
	for i := 0; i < 2; i++ {
		list.toTail(NewNode(i))
	}

	if list.tail != recursiveInverse(list.head, nil) {
		t.Errorf("The head of list must be equal the tail")
	}
}

func BenchmarkReverseNaive(b *testing.B) {
	list := linkedList{}
	for i := 0; i < 1000000; i++ {
		list.toTail(NewNode(i))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		naiveInverse(list.head)
	}
}

func BenchmarkReverseRecursive(b *testing.B) {
	list := linkedList{}
	for i := 0; i < 1000000; i++ {
		list.toTail(NewNode(i))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		recursiveInverse(list.head, nil)
	}
}
