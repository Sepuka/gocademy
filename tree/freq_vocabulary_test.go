package tree

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestFrequencyVocabulary(t *testing.T) {
	var testCases = map[string]struct{
		lst []string
		expected *frequencyVocabularyNode
	}{
		"empty": {
			lst:[]string{},
			expected:nil,
		},
		"some items": {
			lst: []string{"1","2","3","2","1"},
			expected: &frequencyVocabularyNode{
				value: "1",
				cnt:   2,
				right: &frequencyVocabularyNode{
					value: "2",
					cnt:   2,
					right: &frequencyVocabularyNode{
						value: "3",
						cnt:   1,
					},
				},
			},
		},
	}
	var dict *frequencyVocabularyNode
	for nameTest, suiteTest := range testCases {
		for _, word := range suiteTest.lst {
			dict = insert(dict, word)
		}
		assert.Equal(t, dict, suiteTest.expected, fmt.Sprintf("test '%s' failed", nameTest))
	}
}
