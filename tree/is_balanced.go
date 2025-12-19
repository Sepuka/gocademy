package tree

import (
	"math"
)

func isBalanced(root *Node) bool {
	if root == nil {
		return true
	}

	l := getDepth(root.Left)
	if l == -1 {
		return false
	}

	r := getDepth(root.Right)
	if r == -1 {
		return false
	}

	return math.Abs(float64(l-r)) < 2
}

func getDepth(root *Node) int {
	if root == nil {
		return 0
	}

	l := getDepth(root.Left)
	r := getDepth(root.Right)

	if l == -1 || r == -1 {
		return -1
	}

	if math.Abs(float64(l-r)) > 1 {
		return -1
	}

	return max(l, r) + 1
}
