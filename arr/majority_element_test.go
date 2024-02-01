package arr

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMajorityElement(t *testing.T) {
	var testCases = []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{},
			expected: 0,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
		{
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, majorityElement(testCase.nums), testCase.expected)
	}
}
