package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeLen(count, length int) {
	for i := 0; i < count; i++ {
		_ = make([]int, length)
	}
}

func BenchmarkMakeLen_0(b *testing.B) {
	benchmarkMakeLen(b.N, 0)
}

func BenchmarkMakeLen_1(b *testing.B) {
	benchmarkMakeLen(b.N, 1)
}

func BenchmarkMakeLen_2(b *testing.B) {
	benchmarkMakeLen(b.N, 2)
}

func BenchmarkMakeLen_4(b *testing.B) {
	benchmarkMakeLen(b.N, 4)
}

func BenchmarkMakeLen_8(b *testing.B) {
	benchmarkMakeLen(b.N, 8)
}

func BenchmarkMakeLen_16(b *testing.B) {
	benchmarkMakeLen(b.N, 16)
}

func BenchmarkMakeLen_32(b *testing.B) {
	benchmarkMakeLen(b.N, 32)
}

func BenchmarkMakeLen_64(b *testing.B) {
	benchmarkMakeLen(b.N, 64)
}

func BenchmarkMakeLen_128(b *testing.B) {
	benchmarkMakeLen(b.N, 128)
}

func BenchmarkMakeLen_256(b *testing.B) {
	benchmarkMakeLen(b.N, 256)
}

func BenchmarkMakeLen_512(b *testing.B) {
	benchmarkMakeLen(b.N, 512)
}

func BenchmarkMakeLen_1024(b *testing.B) {
	benchmarkMakeLen(b.N, 1024)
}

func BenchmarkMakeLen_2048(b *testing.B) {
	benchmarkMakeLen(b.N, 2048)
}

func BenchmarkMakeLen_4096(b *testing.B) {
	benchmarkMakeLen(b.N, 4096)
}

func BenchmarkMakeLen_8192(b *testing.B) {
	benchmarkMakeLen(b.N, 8192)
}
