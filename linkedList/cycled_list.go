package linkedList

func hasCycle(head *Node) bool {
	for head.next != nil {
		if head.key < 0 {
			return true
		}
		head.key = -1
		head = head.next
	}

	return false
}
