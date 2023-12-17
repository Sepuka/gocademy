package arr

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_singleNumber(t *testing.T) {
	var testCases = []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{2, 2, 3, 2},
			expected: 3,
		},
		{
			numbers:  []int{0, 1, 0, 1, 0, 1, 99},
			expected: 99,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, singleNumber(testCase.numbers), testCase.expected)
	}
}
