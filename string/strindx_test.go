package string

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestStrStr(t *testing.T) {
	var cases = []struct {
		haystack string
		needle   string
		pos      int
	}{
		{
			haystack: "sadbutsad",
			needle:   "sad",
			pos:      0,
		},
		{
			haystack: "sad",
			needle:   "sad",
			pos:      0,
		},
		{
			haystack: "sa",
			needle:   "sad",
			pos:      -1,
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
			pos:      -1,
		},
		{
			haystack: "mississippi",
			needle:   "issip",
			pos:      4,
		},
		{
		  haystack: "mississippi",
		  needle:   "issipi",
		  pos:      -1,
		},
	}

	for _, set := range cases {
		assert.Equal(t, strStr(set.haystack, set.needle), set.pos)
	}
}
