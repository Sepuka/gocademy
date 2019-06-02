package tree

import (
	"fmt"
	"strings"
)

const absentChild = "X"

type ViewItem struct {
	value string
	offset uint
}

func (v *ViewItem) String() string {
	return fmt.Sprintf("%s%s", strings.Repeat(" ", int(v.offset)), v.value)
}

type View struct {
	rows [][]*ViewItem
}

func NewView(node *Node) View {
	view := View{}
	view.rows = make([][]*ViewItem, node.getDeep())
	top := []*ViewItem{
		view.fill(node, 1),
	}
	view.rows[0] = top

	return view
}

func (view *View) fill(node *Node, deep uint) *ViewItem {
	var (
		offset   uint
		left, right, wall *ViewItem
	)

	if node.isLeaf() {
		return &ViewItem{
			value: fmt.Sprintf("%d", node.Value),
			offset: 0,
		}
	}

	left = view.fill(node.Left, deep + 1)
	offset = left.offset * 2 + 1
	if node.Right == nil {
		right = &ViewItem{
			value: absentChild,
			offset: offset,
		}
	} else {
		right = view.fill(node.Right, deep + 1)
	}
	right.offset = offset

	if len(view.rows[deep]) > 0{
		wall = &ViewItem{value: "", offset:offset - left.offset}
		view.rows[deep] = append(view.rows[deep], wall)
	}

	view.rows[deep] = append(view.rows[deep], left, right)

	return &ViewItem{
		value: fmt.Sprintf("%d", node.Value),
		offset: offset,
	}
}

func (view *View) Items() [][]*ViewItem {
	return view.rows
}
