package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	var testCases = []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "b",
			t:        "s",
			expected: false,
		},
		{
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, isSubsequence(testCase.s, testCase.t), fmt.Sprintf("case with '%s' failed", testCase.s))
	}
}
