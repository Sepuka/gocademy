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
			input:    [][]int{{2, 6}, {1, 3}, {8, 10}, {12, 16}},
			expected: [][]int{{1, 6}, {8, 10}, {12, 16}},
		},
		{
			input:    [][]int{{2, 6}, {3, 4}, {8, 10}, {12, 16}},
			expected: [][]int{{2, 6}, {8, 10}, {12, 16}},
		},
		{
			input:    [][]int{{2, 6}, {1, 7}, {8, 10}, {12, 16}},
			expected: [][]int{{1, 7}, {8, 10}, {12, 16}},
		},
		{
			input:    [][]int{{2, 6}, {8, 10}, {12, 16}, {1, 7}},
			expected: [][]int{{8, 10}, {12, 16}, {1, 7}},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, getIntersection(testCase.input), testCase.expected)
	}
}
