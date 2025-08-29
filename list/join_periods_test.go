package list

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_JoinPeriods(t *testing.T) {
	cases := []struct {
		src []int
		exp string
	}{
		{
			src: []int{},
			exp: "",
		}, {
			src: []int{1, 4, 5, 2, 3, 9, 8, 11, 0, 13},
			exp: "0-5,8-9,11,13",
		},
	}

	for _, testCase := range cases {
		assert.Equal(t, JoinPeriods(testCase.src), testCase.exp)
	}
}
