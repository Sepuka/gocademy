package tree

func compare(p *Node, q *Node) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil {
		return false
	}

	if q == nil && p != nil {
		return false
	}

	if p.Value != q.Value {
		return false
	}

	if compare(p.Left, q.Left) == false {
		return false
	}

	if compare(p.Right, q.Right) == false {
		return false
	}

	return true
}
