package search

import (
	"crypto/rand"
	"errors"
	"fmt"
)

func bufferDataProvider(bufferLen uint, patternLen uint) (string, string, error) {
	buffer, err := generateBytes(bufferLen)
	if err != nil {
		return "", "", errors.New(err.Error())
	}

	pattern, err := generateBytes(patternLen)
	if err != nil {
		return "", "", errors.New(err.Error())
	}

	return fmt.Sprint(buffer), fmt.Sprint(pattern), nil
}

func generateBytes(length uint) (n int, err error) {
	buffer := make([]byte, length)

	return rand.Read(buffer)
}
