package tree

import (
	"strings"

)

const (
	wordsDelimiter = " "
	rowsDelimiter = "\n"
)

func BuildFreqVocabulary(text string) *FrequencyVocabularyNode {
	var dict *FrequencyVocabularyNode

	for num, row := range strings.Split(text, rowsDelimiter) {
		if row == "" {
			continue
		}
		for _, chunk := range strings.Split(row, wordsDelimiter) {
			dict = search(dict, word(chunk), num)
		}
	}

	return dict
}
