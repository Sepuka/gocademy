package linkedList

func naiveInverse(first *node) *node {
	var nextLink, third, tail *node
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

func recursiveInverse(current *node, previous *node) *node {
	next := current.next
	current.next = previous
	if next != nil {
		return recursiveInverse(next, current)
	}

	return current
}