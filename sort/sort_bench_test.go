package sort

import (
	"github.com/sepuka/gonerator/number"
	"testing"
)

const sequenceLen = 84190

var sequence []uint32

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
	sequence = number.Uint32(sequenceLen)
}
