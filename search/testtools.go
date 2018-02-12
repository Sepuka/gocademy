package search

import (
	"crypto/rand"
	"errors"
)

func bufferDataProvider(bufferLen uint, patternLen uint) (string, string, error) {
	buffer := make([] byte, bufferLen)
	_, err := rand.Read(buffer)
	pattern := make([] byte, patternLen)
	_, symbolErr := rand.Read(pattern)
	if err == nil && symbolErr == nil {
		return string(buffer), string(pattern), nil
	}

	return "", "", errors.New("Cannot generate test data")
}
