package linkedList

func deleteDuplicates(head *Node) *Node {
	var last = head
	var cur = head.next

	for cur != nil {
		if cur.key != last.key {
			last.next = cur
			last = last.next
		}

		cur = cur.next
	}

	last.next = nil

	return head
}
