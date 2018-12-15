package search

import (
	"errors"
)

func LinearSearch(data string, symbol string) (result int, err error) {
	for index := range data {
		if symbol == string(data[index]) {
			return index, nil
		}
	}

	return -1, errors.New("Symbol not found")
}

// Упрощенный инвариант цикла работает гораздо быстрее
func LinearImprovedSearch(data string, symbol string) (result int, err error) {
	data += symbol
	i := 0
	for data[i] != byte(symbol[0]) {
		i++
	}

	if i+1 == len(data) {
		return -1, errors.New("Symbol not found")
	} else {
		return i, nil
	}
}
