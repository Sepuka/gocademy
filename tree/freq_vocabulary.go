package tree

import (
	"fmt"
	"io"
	"strings"
)

type word string

type Item struct {
	rowNum int
	next *Item
}

type FrequencyVocabularyNode struct {
	value word
	cnt uint
	first, last *Item
	left, right *FrequencyVocabularyNode
}

func Print(f *FrequencyVocabularyNode, writer io.Writer) {
	if f == nil {
		return
	}

	Print(f.left, writer)

	_, _ = writer.Write([]byte(f.value))
	_, _ = writer.Write([]byte("\t"))

	var rows = []string{""}
	item := f.first
	for item != nil {
		rows = append(rows, fmt.Sprintf("%d", item.rowNum))
		item = item.next
	}
	_, _ = writer.Write([]byte(strings.Join(rows, "\t")))
	_, _ = writer.Write([]byte("\n"))

	Print(f.right, writer)
}

func search(node *FrequencyVocabularyNode, word word, num int) *FrequencyVocabularyNode {
	if node == nil {
		item := &Item{
			rowNum: num,
		}
		node = &FrequencyVocabularyNode{
			value: word,
			cnt:   1,
			first: item,
			last: item,
		}
	} else if word < node.value {
		node.left = search(node.left, word, num)
	} else if word > node.value {
		node.right = search(node.right, word, num)
	} else {
		item := &Item{
			rowNum: num,
		}
		node.last.next = item
		node.last = item
	}

	return node
}
