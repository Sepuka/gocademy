package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type viewTest struct {
	suite.Suite
}

func (v *viewTest) TestViewItemToString() {
	cases := []struct {
		viewItem ViewItem
		row string
	}{
		{
			viewItem: ViewItem{
				value: "1",
				offset: 0,
			},
			row: "1",
		},{
			viewItem:ViewItem{
				value: "2",
				offset: 5,
			},
			row: "     2",
		},
	}

	for _, testCase := range cases {
		v.Suite.Equal(testCase.row, testCase.viewItem.String())
	}
}

func (v *viewTest) TestViewItems() {
	cases := []struct {
		node *Node
		items [][]*ViewItem
	}{
		{
			node: &Node{},
			items: [][]*ViewItem{
				{
					&ViewItem{
						value: "0",
					},
				},
			},
		},
		{
			node: &Node{
				Value: 1,
			},
			items: [][]*ViewItem{
				{
					&ViewItem{
						value:  "1",
						offset: 0,
					},
				},
			},
		},
		{
			node: &Node{
				Value: 1,
				Left: &Node{
					Value: 2,
				},
				Right: &Node{
					Value: 3,
				},
			},
			items: [][]*ViewItem{
				{
					&ViewItem{
						value:  "1",
						offset: 1,
					},
				},
				{
					&ViewItem{
						value:  "2",
						offset: 0,
					},
					&ViewItem{
						value:  "3",
						offset: 1,
					},
				},
			},
		},
	}

	for _, testCase := range cases {
		view := NewView(testCase.node)
		i := view.Items()
		v.Suite.Equal(testCase.items, i)
	}
}

func (v *viewTest) TestViewTree() {
	const expectedTree = `
       1
   2       2
 3   3   3   3
4 4 4 4 4 4 4 4`
	var actualTree = ""
	builder := NewTreeBuilder("1 2 3 4 4 3 4 4 2 3 4 4 3 4 4")
	view := NewView(builder.Build(builder.Length()))
	for _, item := range view.Items() {
		row := ""
		for _, r := range item {
			row += r.String()
		}
		actualTree = fmt.Sprintf("%s\n%s", actualTree, row)
	}
	v.Suite.Equal(expectedTree, actualTree)
}

func TestView(t *testing.T) {
	suite.Run(t, new(viewTest))
}
