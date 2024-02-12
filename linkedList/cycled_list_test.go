package linkedList

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestHasCycle(t *testing.T) {
	var testCases = []struct {
		data *Node
		exp  bool
	}{
		{
			data: &Node{},
			exp:  false,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, hasCycle(testCase.data), testCase.exp)
	}
}
