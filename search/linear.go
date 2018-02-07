package linear

import "errors"

func LinearSearch(data string, symbol string) (result int, err error) {
	for current := range data {
		if symbol == string(data[current]) {
			return current, nil
		}
	}

	return -1, errors.New("Pattern not found")
}
