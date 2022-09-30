package linkedList

import (
	"strconv"
)

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	var str, digit string
	var sum, result, surplus int

	for l1 != nil && l2 != nil {
		sum = l1.key + l2.key + surplus
		if sum >= 10 {
			result = sum - 10
		} else {
			result = sum
		}
		digit = strconv.FormatInt(int64(result), 10)
		if sum >= 10 {
			surplus = 1
		} else {
			surplus = 0
		}
		str = digit + str
		l1 = l1.next
		l2 = l2.next
	}

	if l2 != nil {
		l1 = l2
	}
	for l1 != nil {
		result = l1.key + surplus
		surplus = 0
		if result == 10 {
			if l1.next == nil {
				digit = `10`
			} else {
				digit = `0`
				surplus = 1
			}
			str = digit + str
		} else {
			surplus = 0
			digit = strconv.FormatInt(int64(result), 10)
			str = digit + str
		}
		l1 = l1.next
	}

	if surplus > 0 {
		str = `1` + str
	}

	var node *Node

	for i := 0; i < len(str); i++ {
		val, _ := strconv.ParseInt(string(str[i]), 10, 32)
		node = &Node{
			key:  int(val),
			next: node,
		}
	}

	return node
}
