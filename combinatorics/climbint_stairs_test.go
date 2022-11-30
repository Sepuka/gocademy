package combinatorics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	var dataSet = []struct {
		steps    int
		expected int
	}{
		{steps: 0, expected: 0},
		{steps: 1, expected: 1},
		{steps: 2, expected: 2},
		{steps: 3, expected: 3},
		{steps: 4, expected: 5},
		{steps: 5, expected: 8},
		{steps: 6, expected: 13},
	}

	for _, i := range dataSet {
		assert.Equal(t, i.expected, climbStairs(i.steps))
	}
}
