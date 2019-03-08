package linkedList

type node struct {
	key  int
	next *node
	data interface{}
}

func NewNode(key int) *node {
	return &node{key: key}
}

type linkedList struct {
	head *node
	tail *node
	len  int
}

func (l *linkedList) list() []*node {
	result := make([]*node, l.len, l.len)
	element := l.head
	for i := 0; i < l.len; i++ {
		result[i] = element
		element = element.next
	}

	return result
}

func (l *linkedList) toHead(node *node) {
	node.next = l.head
	l.head = node
	l.len += 1
}