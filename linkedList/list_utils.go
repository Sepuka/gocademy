package linkedList

func naiveInverse(first *Node) *Node {
	var nextLink, third, tail *Node
	tail = first
	nextLink = first.next
	first.next = nil
	for nextLink != nil {
		third = nextLink.next
		nextLink.next = first
		tail = nextLink
		nextLink = third
	}

	return tail
}

func recursiveInverse(current *Node, previous *Node) *Node {
	next := current.next
	current.next = previous
	if next != nil {
		return recursiveInverse(next, current)
	}

	return current
}