package arr

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMerge(t *testing.T) {
	var testCases = []struct {
		num1     []int
		n        int
		num2     []int
		m        int
		expected []int
	}{
		{
			num1:     []int{0},
			n:        0,
			num2:     []int{1},
			m:        1,
			expected: []int{1},
		},
		{
			num1:     []int{1},
			n:        1,
			num2:     []int{},
			m:        0,
			expected: []int{1},
		},
		{
			num1:     []int{1, 2, 3, 0, 0, 0},
			n:        3,
			num2:     []int{2, 5, 6},
			m:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
	}

	for _, testCase := range testCases {
		merge(testCase.num1, testCase.n, testCase.num2, testCase.m)
		assert.Equal(t, testCase.num1, testCase.expected)
	}
}
