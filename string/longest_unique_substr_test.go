package string

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestLongestSubstr(t *testing.T) {
	var dataSet = []struct {
		input  string
		output int
	}{
		{
			input:  "",
			output: 0,
		},
    {
      input:  "a",
      output: 1,
    },
    {
      input:  "ab",
      output: 2,
    },
    {
      input:  "abcabc",
      output: 3,
    },
  }

  for _, set := range dataSet {
    assert.Equal(t, lengthOfLongestSubstring(set.input), set.output)
  }
}
