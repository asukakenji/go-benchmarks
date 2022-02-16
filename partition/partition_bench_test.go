package partition_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/partition"
)

// Exported (global) variable serving as input for some
// of the benchmarks to ensure side-effect free calls
// are not optimized away.
var Input int

// Exported (global) variable to store function results
// during benchmarking to ensure side-effect free calls
// are not optimized away.
var Output int

func benchmarkPartition(b *testing.B, partition func(int) [][]int, n int) {
	var s int
	for i := 0; i < b.N; i++ {
		s += partition((i*97)%(n-1) + 1)[0][0]
	}
	Output = s
}

func BenchmarkPartition1_4(b *testing.B) {
	benchmarkPartition(b, partition.Partition1, 4)
}

func BenchmarkPartition1_8(b *testing.B) {
	benchmarkPartition(b, partition.Partition1, 8)
}

func BenchmarkPartition1_16(b *testing.B) {
	benchmarkPartition(b, partition.Partition1, 16)
}

func BenchmarkPartition1_32(b *testing.B) {
	benchmarkPartition(b, partition.Partition1, 32)
}

func BenchmarkPartition1_48(b *testing.B) {
	benchmarkPartition(b, partition.Partition1, 48)
}

// ---

func BenchmarkPartition2_4(b *testing.B) {
	benchmarkPartition(b, partition.Partition2, 4)
}

func BenchmarkPartition2_8(b *testing.B) {
	benchmarkPartition(b, partition.Partition2, 8)
}

func BenchmarkPartition2_16(b *testing.B) {
	benchmarkPartition(b, partition.Partition2, 16)
}

func BenchmarkPartition2_32(b *testing.B) {
	benchmarkPartition(b, partition.Partition2, 32)
}

func BenchmarkPartition2_48(b *testing.B) {
	benchmarkPartition(b, partition.Partition2, 48)
}

// ---

func BenchmarkPartition2a_4(b *testing.B) {
	benchmarkPartition(b, partition.Partition2a, 4)
}

func BenchmarkPartition2a_8(b *testing.B) {
	benchmarkPartition(b, partition.Partition2a, 8)
}

func BenchmarkPartition2a_16(b *testing.B) {
	benchmarkPartition(b, partition.Partition2a, 16)
}

func BenchmarkPartition2a_32(b *testing.B) {
	benchmarkPartition(b, partition.Partition2a, 32)
}

func BenchmarkPartition2a_48(b *testing.B) {
	benchmarkPartition(b, partition.Partition2a, 48)
}

// ---

func BenchmarkPartition3_4(b *testing.B) {
	benchmarkPartition(b, partition.Partition3, 4)
}

func BenchmarkPartition3_8(b *testing.B) {
	benchmarkPartition(b, partition.Partition3, 8)
}

func BenchmarkPartition3_16(b *testing.B) {
	benchmarkPartition(b, partition.Partition3, 16)
}

func BenchmarkPartition3_32(b *testing.B) {
	benchmarkPartition(b, partition.Partition3, 32)
}

func BenchmarkPartition3_48(b *testing.B) {
	benchmarkPartition(b, partition.Partition3, 48)
}
