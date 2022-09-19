package linkedList

type Node struct {
	key  int
	next *Node
	data interface{}
}

func NewNode(key int) *Node {
	return &Node{key: key}
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func (l *LinkedList) Set(key int, value interface{}, len int) int {
	var (
		cur  = l.head
		prev *Node
		node = NewNode(key)
	)

	node.data = value

	if l.len == 0 {
		l.addToHead(node)
		return 0
	}

	if cur.key == key {
		return 0
	}

	for cur = cur.next; cur != nil && cur.next != nil; cur = cur.next {
		if cur.key == key {
			l.addToHead(cur)
			prev.next = cur.next
			return 0
		}
		prev = cur
	}

	if l.len < len {
		l.addToHead(node)
		return 0
	}

	if prev != nil {
		prev.next = nil
	}
	l.addToHead(node)

	return cur.key

}

func (l *LinkedList) list() []*Node {
	result := make([]*Node, l.len, l.len)
	element := l.head
	for i := 0; i < l.len; i++ {
		result[i] = element
		element = element.next
	}

	return result
}

func (l *LinkedList) addToHead(node *Node) {
	node.next = l.head
	l.head = node
	l.len++
}

func (l *LinkedList) addToTail(node *Node) {
	if l.tail == nil {
		if l.head == nil {
			l.addToHead(node)
			l.tail = node
		} else {
			tail := l.head
			for tail.next != nil {
				tail = tail.next
			}
			tail.next = node
			l.len++
		}
	} else {
		l.tail.next = node
		l.tail = node
		l.len++
	}
}

func (l *LinkedList) toPos(nodeDescr *Node, pos int) {
	if pos > l.len {
		panic("undefined position")
	}
	if pos == 0 {
		l.addToHead(nodeDescr)
		return
	}
	if pos == l.len {
		l.addToTail(nodeDescr)
		return
	}

	var current, buffer *Node
	current = l.head
	cnt := 0
	for cnt < pos {
		cnt++
		current = current.next
	}

	buffer = &Node{key: current.key, next: current.next, data: current.data}
	current.key = nodeDescr.key
	current.data = nodeDescr.data
	current.next = buffer
	l.len++
}

func (head *Node) middle() *Node {
	if head == nil {
		return nil
	}

	var next = head.next
	if next == nil {
		return head
	}

	var middle = next
	var length = 2
	var pos = 2

	for next.next != nil {
		next = next.next
		length++
		if length/2 >= pos {
			middle = middle.next
			pos++
		}
	}

	return middle
}
