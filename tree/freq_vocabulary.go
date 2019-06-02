package tree

type frequencyVocabularyNode struct {
	value int
	cnt uint
	left *frequencyVocabularyNode
	right *frequencyVocabularyNode
}

func insert(node *frequencyVocabularyNode, x int) *frequencyVocabularyNode {
	if node == nil {
		node = &frequencyVocabularyNode{
			value: x,
			cnt:   1,
		}
	} else if x < node.value {
		node.left = insert(node.left, x)
	} else if x > node.value {
		node.right = insert(node.right, x)
	} else {
		node.cnt++
	}

	return node
}
