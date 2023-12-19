package arr

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_removeDuplicates(t *testing.T) {
	var testCases = []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 1, 2},
			expected: 2,
		},
		{
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
		{
			nums:     []int{1, 2},
			expected: 2,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, removeDuplicates(testCase.nums), testCase.expected)
	}
}
