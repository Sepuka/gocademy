package sort

import (
	"github.com/sepuka/gonerator/number"
	"testing"
)

const sequenceLen = 50000

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

func BenchmarkImprovedBubbleSortBench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		sequence = number.SeqUint32(sequenceLen)
		b.StartTimer()
		bubble(sequence)
	}
}

func BenchmarkShellSortBench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		sequence = number.SeqUint32(sequenceLen)
		b.StartTimer()
		shell(sequence)
	}
}

func BenchmarkHeapSortBench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		sequence = number.SeqUint32(sequenceLen)
		b.StartTimer()
		heap(sequence)
	}
}
