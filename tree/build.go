package tree

import (
	"strconv"
	"strings"
)

const separator = " "

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) isLeaf() bool {
	return node.Left == nil
}

func (node *Node) getDeep() uint {
	if node == nil {
		return 0
	}

	return calcDeep(node, 1)
}

func calcDeep(node *Node, deep uint) uint {
	if node.isLeaf() {
		return deep
	}

	return calcDeep(node.Left, deep+1)
}

type Builder struct {
	source []string
}

func NewTreeBuilder(source string) Builder {
	return Builder{source: strings.Split(source, separator)}
}

func (t *Builder) Length() int {
	return len(t.source)
}

func (t *Builder) Build(n int) *Node {
	var (
		node = &Node{}
		value int64
		err error
		nl, nr int
	)
	if n == 0 {
		return nil
	} else {
		value, err = strconv.ParseInt(t.source[0], 10, 32)
		if err != nil {
			return nil
		}
		node.Value, t.source = int(value), t.source[1:]
		nl = n / 2
		nr = n - nl - 1
		node.Left = t.Build(nl)
		node.Right = t.Build(nr)
	}

	return node
}
