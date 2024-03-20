package arr

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetIntersection(t *testing.T) {
	var testCases = []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{[]int{2, 6}, []int{1, 3}, []int{8, 10}, []int{12, 16}},
			expected: [][]int{[]int{1, 6}, []int{8, 10}, []int{12, 16}},
		},
		{
			input:    [][]int{[]int{2, 6}, []int{3, 4}, []int{8, 10}, []int{12, 16}},
			expected: [][]int{[]int{2, 6}, []int{8, 10}, []int{12, 16}},
		},
		{
			input:    [][]int{[]int{2, 6}, []int{1, 7}, []int{8, 10}, []int{12, 16}},
			expected: [][]int{[]int{1, 7}, []int{8, 10}, []int{12, 16}},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, getIntersection(testCase.input), testCase.expected)
	}
}
