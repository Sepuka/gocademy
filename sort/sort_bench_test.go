package sort

import (
	"github.com/sepuka/gocademy/test"
	"testing"
)

const sequenceLen = 84190

var sequence []int

func BenchmarkDummySortBench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy(sequence)
	}
}

func BenchmarkImprovedDummySortBench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		improvedDummy(sequence)
	}
}

func init() {
	sequence = test.GenerateIntSequence(sequenceLen)
}
