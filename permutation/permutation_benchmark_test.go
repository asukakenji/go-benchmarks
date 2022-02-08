package permutation_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation"
)

func dummySliceConsumer[T any]([]T) {}

func benchmarkPermutationOrder0[T any](b *testing.B, size int) {
	slice := make([]T, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		permutation.PermutationOrder0(slice, dummySliceConsumer[T])
	}
}

func benchmarkPermutationOrder1[T any](b *testing.B, size int) {
	slice := make([]T, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		permutation.PermutationOrder1(dummySliceConsumer[T], slice)
	}
}

func BenchmarkPermutationOrder0_8(b *testing.B) {
	benchmarkPermutationOrder0[uint](b, 8)
}

func BenchmarkPermutationOrder0_9(b *testing.B) {
	benchmarkPermutationOrder0[uint](b, 9)
}

func BenchmarkPermutationOrder0_10(b *testing.B) {
	benchmarkPermutationOrder0[uint](b, 10)
}

func BenchmarkPermutationOrder1_8(b *testing.B) {
	benchmarkPermutationOrder1[uint](b, 8)
}

func BenchmarkPermutationOrder1_9(b *testing.B) {
	benchmarkPermutationOrder1[uint](b, 9)
}

func BenchmarkPermutationOrder1_10(b *testing.B) {
	benchmarkPermutationOrder1[uint](b, 10)
}
