package tree

import (
	"strconv"
	"strings"
)

type treeNode struct {
	value int
	left *treeNode
	right *treeNode
}

type Builder struct {
	source []string
}

func NewTreeBuilder(source string) Builder {
	return Builder{source: strings.Split(source, " ")}
}

func (t *Builder) build(n int) *treeNode {
	var (
		node = &treeNode{}
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
		node.value, t.source = int(value), t.source[1:]
		nl = n / 2
		nr = n - nl - 1
		node.left = t.build(nl)
		node.right = t.build(nr)
	}

	return node
}
