package list

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_expectedSum(t *testing.T) {
	cases := []struct {
		l []int
		s int
		e bool
	}{
		{
			l: []int{},
			s: 0,
			e: false,
		},
		{
			l: []int{10, 15, 3, 7},
			s: 17,
			e: true,
		},
		{
			l: []int{1, 2, 3, 4, 5, 6},
			s: 12,
			e: false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.e, ContainsSum(c.l, c.s))
		})
	}
}
