package tree

import "testing"

func TestF(t *testing.T) {
	var dict = &frequencyVocabularyNode{}
	lst := []int{1,2,3,4,5,6}
	for _, word := range lst {
		dict = insert(dict, word)
	}
}
