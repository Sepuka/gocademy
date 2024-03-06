package numeral_system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRomanRecur(t *testing.T) {
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
		assert.Equal(t, testCase.output, intToRomanRecur(testCase.input))
	}
}

func TestIntToRomanIter(t *testing.T) {
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
		assert.Equal(t, testCase.output, intToRomanIter(testCase.input))
	}
}

func TestIntToRomanModule(t *testing.T) {
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
		assert.Equal(t, testCase.output, intToRomanModule(testCase.input))
	}
}

func BenchmarkIntToRomanRecur(t *testing.B) {
	for i := 0; i < 9999; i++ {
		intToRomanRecur(i)
	}
}

func BenchmarkIntToRomanIter(t *testing.B) {
	for i := 0; i < 9999; i++ {
		intToRomanIter(i)
	}
}

func BenchmarkIntToRomanModule(t *testing.B) {
	for i := 0; i < 9999; i++ {
		intToRomanModule(i)
	}
}
