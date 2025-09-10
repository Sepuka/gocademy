package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JoinSortedPeriods(t *testing.T) {
	cases := []struct {
		src []int
		exp []string
	}{
		{
			src: []int{},
			exp: []string{},
		},
		{
			src: []int{0, 1, 2, 4, 5, 7},
			exp: []string{"0->2", "4->5", "7"},
		},
		{
			src: []int{0, 2, 3, 4, 6, 8, 9},
			exp: []string{"0", "2->4", "6", "8->9"},
		},
	}

	for _, testCase := range cases {
		assert.Equal(t, testCase.exp, summaryRanges(testCase.src))
	}
}
