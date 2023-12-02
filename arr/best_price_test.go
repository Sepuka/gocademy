package arr

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMaxProfit(t *testing.T) {
	var testCases = []struct {
		calendar []int
		expected int
	}{
		{
			calendar: []int{0, 1},
			expected: 1,
		},
		{
			calendar: []int{0, 1, 2},
			expected: 2,
		},
		{
			calendar: []int{0, 1, 2, 1},
			expected: 2,
		},
		{
			calendar: []int{0, 1, 2, 1, 0},
			expected: 2,
		},
		{
			calendar: []int{0, 1, 2, 1, 0, 5},
			expected: 5,
		},
		{
			calendar: []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			calendar: []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, maxProfit(testCase.calendar), testCase.expected)
	}
}
