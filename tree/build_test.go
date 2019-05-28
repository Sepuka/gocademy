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
		tree *treeNode
	}{
		"empty source": {
			source: "",
			tree: nil,
		},
		"alone symbol": {
			source: "1",
			tree: &treeNode{
				value: 1,
			},
		},
		"couple of symbols": {
			source: "1 2",
			tree: &treeNode{
				value: 1,
				left: &treeNode{
					value: 2,
				},
			},
		},
		"triple of symbols": {
			source: "1 2 3",
			tree: &treeNode{
				value: 1,
				left: &treeNode{
					value: 2,
				},
				right: &treeNode{
					value: 3,
				},
			},
		},
		"a lot of symbols": {
			source: "1 2 3 4 5 6 7",
			tree: &treeNode{
				value: 1,
				left: &treeNode{
					value: 2,
					left: &treeNode{
						value: 3,
					},
					right: &treeNode{
						value: 4,
					},
				},
				right: &treeNode{
					value: 5,
					left: &treeNode{
						value: 6,
					},
					right: &treeNode{
						value: 7,
					},
				},
			},
		},
	}
	for testName, testCase := range cases {
		builder := NewTreeBuilder(testCase.source)
		tree := builder.build(len(builder.source))
		t.Suite.Equal(testCase.tree, tree, fmt.Sprintf("case '%s' failed", testName))
	}
}

func Test(t *testing.T) {
	suite.Run(t, new(treeBuilderTest))
}