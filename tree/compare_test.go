package tree

import (
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type (
	testCases struct {
		suite.Suite
		dataSet []struct {
			p        *Node
			q        *Node
			expected bool
		}
	}
)

func (suite *testCases) SetupTest() {
	suite.dataSet = []struct {
		p        *Node
		q        *Node
		expected bool
	}{
		{p: nil, q: nil, expected: true},
		{p: &Node{Value: 1}, q: nil, expected: false},
		{p: &Node{Value: 1}, q: &Node{Value: 1}, expected: true},
		{p: &Node{Value: 1, Left: &Node{Value: 2}, Right: &Node{Value: 3}}, q: &Node{Value: 1, Left: &Node{Value: 2}, Right: &Node{Value: 3}}, expected: true},
		{p: &Node{Value: 1, Left: &Node{Value: 2}, Right: &Node{Value: 3}}, q: &Node{Value: 1, Left: &Node{Value: 2}, Right: &Node{Value: 2}}, expected: false},
	}
}

func (suite *testCases) Test() {
	for _, set := range suite.dataSet {
		assert.Equal(suite.T(), compare(set.p, set.q), set.expected)
	}
}

func TestCompare(t *testing.T) {
	suite.Run(t, new(testCases))
}
