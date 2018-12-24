package sort

import (
	"github.com/sepuka/gonerator/number"
	"testing"
)

const sequenceLen = 100000

var sequence []uint32

func BenchmarkDummySortBench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		sequence = number.SeqUint32(sequenceLen)
		b.StartTimer()
		dummy(sequence)
	}
}

func BenchmarkImprovedDummySortBench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		sequence = number.SeqUint32(sequenceLen)
		b.StartTimer()
		improvedDummy(sequence)
	}
}
