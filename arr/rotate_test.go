package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	var testCases = []struct {
		arr      []int
		k        int
		expected []int
	}{
		{
			arr:      []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			arr:      []int{-1},
			k:        2,
			expected: []int{-1},
		},
		{
			arr:      []int{1, 2},
			k:        3,
			expected: []int{2, 1},
		},
		{
			arr:      []int{1, 2},
			k:        2,
			expected: []int{1, 2},
		},
		{
			arr:      []int{1, 2, 3},
			k:        4,
			expected: []int{3, 1, 2},
		},
	}

	for _, testCase := range testCases {
		rotate(testCase.arr, testCase.k)
		assert.Equal(t, testCase.expected, testCase.arr)
	}
}
