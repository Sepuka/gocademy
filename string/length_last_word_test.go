package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	var testCases = []struct {
		input    string
		expected int
	}{
		{
			input:    "Hello world",
			expected: 5,
		},
		{
			input:    "luffy is still joyboy",
			expected: 6,
		},
		{
			input:    "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			input:    "a",
			expected: 1,
		},
		{
			input:    "a ",
			expected: 1,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, lengthOfLastWord(testCase.input), fmt.Sprintf("case with '%s' failed", testCase.input))
	}
}
