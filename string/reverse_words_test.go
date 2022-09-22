package string

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	var dataSet = []struct {
		input  string
		output string
	}{
		{
			input:  "",
			output: "",
		},
		{
			input:  "a",
			output: "a",
		},
		{
			input:  "aa",
			output: "aa",
		},
		{
			input:  "ab",
			output: "ba",
		},
		{
			input:  "mama",
			output: "amam",
		},
		{
			input:  "hello world",
			output: "olleh dlrow",
		},
	}

	for _, set := range dataSet {
		assert.Equal(t, set.output, reverseWords(set.input))
	}
}
