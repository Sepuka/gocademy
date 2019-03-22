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

func (l *linkedList) toHead(nodeDescr *node) {
	nodeDescr.next = l.head
	l.head = nodeDescr
	l.len++
}

func (l *linkedList) toTail(nodeDescr *node) {
	if l.tail == nil {
		if l.head == nil {
			l.toHead(nodeDescr)
			l.tail = nodeDescr
		} else {
			tail := l.head
			for tail.next != nil {
				tail = tail.next
			}
			tail.next = nodeDescr
			l.len++
		}
	} else {
		l.tail.next = nodeDescr
		l.tail = nodeDescr
		l.len++
	}
}

func (l *linkedList) toPos(nodeDescr *node, pos int) {
	if pos > l.len {
		panic("undefined position")
	}
	if pos == 0 {
		l.toHead(nodeDescr)
		return
	}
	if pos == l.len {
		l.toTail(nodeDescr)
		return
	}

	var current, buffer *node
	current = l.head
	cnt := 0
	for cnt < pos {
		cnt++
		current = current.next
	}

	buffer = &node{key: current.key, next: current.next, data: current.data}
	current.key = nodeDescr.key
	current.data = nodeDescr.data
	current.next = buffer
	l.len++
}
