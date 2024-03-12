package numeral_system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	var cases = []struct {
		s string
		i int
	}{
		{
			s: "I",
			i: 1,
		},
		{
			s: "III",
			i: 3,
		},
		{
			s: "IV",
			i: 4,
		},
		{
			s: "VI",
			i: 6,
		},
		{
			s: "VIII",
			i: 8,
		},
		{
			s: "XIX",
			i: 19,
		},
		{
			s: "LVIII",
			i: 58,
		},
		{
			s: "MCMXCIV",
			i: 1994,
		},
	}

	for _, testCase := range cases {
		assert.Equal(t, testCase.i, romanToInt(testCase.s))
	}
}
