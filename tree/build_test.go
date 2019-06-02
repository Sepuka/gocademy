package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type treeBuilderTest struct {
	suite.Suite
}

func (t *treeBuilderTest) Test() {
	var cases = map[string]struct{
		source string
		tree *Node
		deep uint
	}{
		"empty source": {
			source: "",
			tree: nil,
		},
		"alone symbol": {
			source: "1",
			tree: &Node{
				Value: 1,
			},
			deep: 1,
		},
		"couple of symbols": {
			source: "1 2",
			tree: &Node{
				Value: 1,
				Left: &Node{
					Value: 2,
				},
			},
			deep: 2,
		},
		"triple of symbols": {
			source: "1 2 3",
			tree: &Node{
				Value: 1,
				Left: &Node{
					Value: 2,
				},
				Right: &Node{
					Value: 3,
				},
			},
			deep: 2,
		},
		"a lot of symbols": {
			source: "1 2 3 4 5 6",
			tree: &Node{
				Value: 1,
				Left: &Node{
					Value: 2,
					Left: &Node{
						Value: 3,
					},
					Right: &Node{
						Value: 4,
					},
				},
				Right: &Node{
					Value: 5,
					Left: &Node{
						Value: 6,
					},
				},
			},
			deep: 3,
		},
	}
	for testName, testCase := range cases {
		errMsg := fmt.Sprintf("case '%s' failed", testName)
		builder := NewTreeBuilder(testCase.source)
		tree := builder.Build(len(builder.source))
		t.Suite.Equal(testCase.deep, tree.getDeep(), errMsg)
		t.Suite.Equal(testCase.tree, tree, errMsg)
	}
}

func Test(t *testing.T) {
	suite.Run(t, new(treeBuilderTest))
}