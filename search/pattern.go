package search

import (
	"errors"
)

func SimplePatternSearch(text string, pattern string) (result int, err error) {
	textPos := 0
	textLen := len(text)
	patternLen := len(pattern)
	for (textPos <= textLen - patternLen) && !isMismatchPosition(textPos, pattern, text) {
		textPos++
	}

	if textLen - patternLen < textPos {
		return -1, errors.New("Symbol not found")
	} else {
		return textPos, nil
	}
}

func isMismatchPosition(textPos int, pattern string, text string) (bool) {
	for patternIndex := range pattern {
		if pattern[patternIndex] == text[textPos] {
			textPos++
		} else {
			return false
		}
	}

	return true
}