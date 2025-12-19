package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_is_balanced(t *testing.T) {
	var cases = []struct {
		input  *Node
		expect bool
	}{
		{
			input:  &Node{Value: 1, Left: &Node{Value: 2, Left: &Node{Value: 3, Left: &Node{Value: 4, Left: &Node{Value: 5}}}}, Right: nil},
			expect: false,
		},
	}
	for _, c := range cases {
		assert.Equal(t, c.expect, isBalanced(c.input))
	}
}
