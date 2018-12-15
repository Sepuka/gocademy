package search

import (
	"testing"
)

type dataSet struct {
	data     string
	symbol   string
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
			t.Error("For", set.symbol, "expected", set.position, "position, but got", actualPos)
		}
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	buffer, symbol, err := bufferDataProvider(100, 1)
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		LinearSearch(buffer, symbol)
	}
}

func TestLinearImprovedSearch(t *testing.T) {
	for _, set := range dataProvider {
		var actualPos, _ = LinearImprovedSearch(set.data, set.symbol)
		if actualPos != set.position {
			t.Error("For", set.symbol, "expected", set.position, "position, but got", actualPos)
		}
	}
}

func BenchmarkLinearImprovedSearch(b *testing.B) {
	buffer, symbol, err := bufferDataProvider(100, 1)
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		LinearImprovedSearch(buffer, symbol)
	}
}
