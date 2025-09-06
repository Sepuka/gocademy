package search

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_subsequence(t *testing.T) {
	cases := []struct {
		s []rune
		t []rune
		r bool
	}{
		//{
		//	s: []rune(""),
		//	t: []rune(""),
		//	r: true,
		//},
		//{
		//	s: []rune("abc"),
		//	t: []rune("ahbgdc"),
		//	r: true,
		//},
		{
			s: []rune("axc"),
			t: []rune("ahbgdc"),
			r: false,
		},
	}

	for _, testCase := range cases {
		assert.Equal(t, testCase.r, isSubsequence(testCase.s, testCase.t))
	}
}
