package numeral_system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	var testCases = []struct {
		input  int
		output string
	}{
		{
			input:  3,
			output: `III`,
		},
		{
			input:  40,
			output: `XL`,
		},
		{
			input:  49,
			output: `XLIX`,
		},
		{
			input:  58,
			output: `LVIII`,
		},
		{
			input:  148,
			output: `CXLVIII`,
		},
		{
			input:  701,
			output: `DCCI`,
		},
		{
			input:  999,
			output: `CMXCIX`,
		},
		{
			input:  1994,
			output: `MCMXCIV`,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.output, intToRoman(testCase.input))
	}
}
