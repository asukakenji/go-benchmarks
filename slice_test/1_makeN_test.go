package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeN(b *testing.B, length int) {
	for i := 0; i < b.N; i++ {
		_ = make([]int, length)
	}
}

func BenchmarkMakeN_0(b *testing.B) {
	benchmarkMakeN(b, 0)
}

func BenchmarkMakeN_1(b *testing.B) {
	benchmarkMakeN(b, 1)
}

func BenchmarkMakeN_2(b *testing.B) {
	benchmarkMakeN(b, 2)
}

func BenchmarkMakeN_4(b *testing.B) {
	benchmarkMakeN(b, 4)
}

func BenchmarkMakeN_8(b *testing.B) {
	benchmarkMakeN(b, 8)
}

func BenchmarkMakeN_16(b *testing.B) {
	benchmarkMakeN(b, 16)
}

func BenchmarkMakeN_32(b *testing.B) {
	benchmarkMakeN(b, 32)
}

func BenchmarkMakeN_64(b *testing.B) {
	benchmarkMakeN(b, 64)
}

func BenchmarkMakeN_128(b *testing.B) {
	benchmarkMakeN(b, 128)
}

func BenchmarkMakeN_256(b *testing.B) {
	benchmarkMakeN(b, 256)
}

func BenchmarkMakeN_512(b *testing.B) {
	benchmarkMakeN(b, 512)
}

func BenchmarkMakeN_1024(b *testing.B) {
	benchmarkMakeN(b, 1024)
}

func BenchmarkMakeN_2048(b *testing.B) {
	benchmarkMakeN(b, 2048)
}

func BenchmarkMakeN_4096(b *testing.B) {
	benchmarkMakeN(b, 4096)
}

func BenchmarkMakeN_8192(b *testing.B) {
	benchmarkMakeN(b, 8192)
}
