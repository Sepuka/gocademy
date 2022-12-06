package linkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicates_WhenIsNil(t *testing.T) {
	assert.Nil(t, nil, deleteDuplicates(nil))
}

func TestDeleteDuplicates(t *testing.T) {
	var list = &LinkedList{}
	list.addToTail(NewNode(1))
	list.addToTail(NewNode(1))
	list.addToTail(NewNode(2))
	//list.addToTail(NewNode(3))
	//list.addToTail(NewNode(3))

	var expected = &LinkedList{}
	expected.addToTail(NewNode(1))
	expected.addToTail(NewNode(2))
	//expected.addToTail(NewNode(3))

	var actual = deleteDuplicates(list.head)

	for actual != nil {
		assert.Equal(t, actual.key, expected.head.key)
		actual = actual.next
		expected.head = expected.head.next
	}

	assert.Nil(t, expected.head)
}
