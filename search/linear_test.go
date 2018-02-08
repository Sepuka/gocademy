package search

import (
	"testing"
)

type dataSet struct {
	data string
	symbol string
	position int
}

var dataProvider = []dataSet{
	{"qwerty", "q", 0},
	{"qwerty", "w", 1},
	{"qwerty", "e", 2},
	{"qwerty", "r", 3},
	{"qwerty", "t", 4},
	{"qwerty", "y", 5},
	{"qwerty", "x", -1},
}

func TestLinearSearch(t *testing.T) {
	for _, set := range dataProvider {
		var actualPos, _ = LinearSearch(set.data, set.symbol)
		if actualPos != set.position {
			t.Error("For", set.data, "expected", set.position, "got", actualPos)
		}
	}
}