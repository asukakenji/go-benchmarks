package math_test

import (
	"testing"

	math "github.com/asukakenji/go-benchmarks/math_test"
)

func TestBitCountUInt32(t *testing.T) {
	fs := []func(uint32) int{
		math.BitCountUInt32Java,
		math.BitCountUInt32CallGCC,
		math.BitCountUInt32GCCImpl,
	}
	names := []string{
		"BitCountUInt32Java",
		"BitCountUInt32CallGCC",
		"BitCountUInt32GCCImpl",
	}
	gen := randomIntsInTheFirstPage()
	for i := 0; i < 512; i++ {
		n := uint32(gen())
		expected := math.BitCountUInt32Naive(n)
		for j, f := range fs {
			got := f(n)
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", names[j], n, got, expected)
			}
		}
	}
}

func TestBitCountUInt64(t *testing.T) {
	fs := []func(uint64) int{
		math.BitCountUInt64Java,
		math.BitCountUInt64CallGCC,
		math.BitCountUInt64GCCImpl,
	}
	names := []string{
		"BitCountUInt64Java",
		"BitCountUInt64CallGCC",
		"BitCountUInt64GCCImpl",
	}
	gen := randomIntsInTheFirstPage()
	for i := 0; i < 512; i++ {
		n := uint64(gen())
		expected := math.BitCountUInt64Naive(n)
		for j, f := range fs {
			got := f(n)
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", names[j], n, got, expected)
			}
		}
	}
}

func benchmarkBitCountUInt32(b *testing.B, f func(uint32) int) {
	gen := randomIntsInTheFirstPage()
	for i, count := 0, b.N; i < count; i++ {
		f(uint32(gen()))
	}
}

func benchmarkBitCountUInt64(b *testing.B, f func(uint64) int) {
	gen := randomIntsInTheFirstPage()
	for i, count := 0, b.N; i < count; i++ {
		f(uint64(gen()))
	}
}

func BenchmarkBitCountUInt32Naive(b *testing.B) {
	benchmarkBitCountUInt32(b, math.BitCountUInt32Naive)
}

func BenchmarkBitCountUInt64Naive(b *testing.B) {
	benchmarkBitCountUInt64(b, math.BitCountUInt64Naive)
}

// java.lang.Integer#bitCount
func BenchmarkBitCountUInt32Java(b *testing.B) {
	benchmarkBitCountUInt32(b, math.BitCountUInt32Java)
}

// java.lang.Long#bitCount
func BenchmarkBitCountUInt64Java(b *testing.B) {
	benchmarkBitCountUInt64(b, math.BitCountUInt64Java)
}

func BenchmarkBitCountUInt32CallGCC(b *testing.B) {
	benchmarkBitCountUInt32(b, math.BitCountUInt32CallGCC)
}

func BenchmarkBitCountUInt64CallGCC(b *testing.B) {
	benchmarkBitCountUInt64(b, math.BitCountUInt64CallGCC)
}

func BenchmarkBitCountUInt32GCCImpl(b *testing.B) {
	benchmarkBitCountUInt32(b, math.BitCountUInt32GCCImpl)
}

func BenchmarkBitCountUInt64GCCImpl(b *testing.B) {
	benchmarkBitCountUInt64(b, math.BitCountUInt64GCCImpl)
}
