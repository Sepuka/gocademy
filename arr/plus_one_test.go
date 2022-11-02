package arr

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestOnePlus(t *testing.T) {
	var dataSet = []struct {
		input    []int
		expected []int
	}{
		{input: []int{1}, expected: []int{2}},
		{input: []int{9}, expected: []int{1, 0}},
		{input: []int{1, 1}, expected: []int{1, 2}},
		{input: []int{1, 2, 3}, expected: []int{1, 2, 4}},
		{input: []int{1, 2, 9}, expected: []int{1, 3, 0}},
		{input: []int{4, 3, 2, 1}, expected: []int{4, 3, 2, 2}},
	}

	for _, i := range dataSet {
		assert.Equal(t, onePlus(i.input), i.expected)
	}
}
