package search

import (
	"testing"
)

type patternDataSet struct {
	data     string
	pattern  string
	position int
}

var patternDataProvider = []patternDataSet{
	{"Go is a programming language created at Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson", "Go", 0},
	{"Go is a programming language created at Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson", "language", 20},
	{"Go is a programming language created at Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson", "Thompson", 94},
	{"Go is a programming language created at Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson", "blah-blah", -1},
}

func TestSimplePatternSearch(t *testing.T) {
	for _, set := range patternDataProvider {
		var actualPos, _ = SimplePatternSearch(set.data, set.pattern)
		if actualPos != set.position {
			t.Error("For", set.pattern, "expected", set.position, "position, but got", actualPos)
		}
	}
}

func BenchmarkSimplePatternSearch(b *testing.B) {
	buffer, pattern, err := bufferDataProvider(100, 10)
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		SimplePatternSearch(buffer, pattern)
	}
}