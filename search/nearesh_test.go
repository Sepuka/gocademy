package search

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_nearest(t *testing.T) {
	cases := []struct {
		l   []int
		i   int
		n   int
		exp []int
	}{
		{
			[]int{2, 3, 5, 7, 11},
			3,
			2,
			[]int{5, 7},
		},
		{
			[]int{4, 12, 15, 15, 24},
			1,
			3,
			[]int{12, 15, 15},
		},
		{
			[]int{2, 3, 5, 7, 11},
			2,
			2,
			[]int{3, 5},
		},
	}

	for _, c := range cases {
		assert.Equal(t, nearest(c.l, c.i, c.n), c.exp)
	}
}
