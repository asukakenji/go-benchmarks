package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeLen_0(b *testing.B) {
	benchmarkMakeLen(b, 0)
}

func BenchmarkMakeLen_1(b *testing.B) {
	benchmarkMakeLen(b, 1)
}

func BenchmarkMakeLen_2(b *testing.B) {
	benchmarkMakeLen(b, 2)
}

func BenchmarkMakeLen_4(b *testing.B) {
	benchmarkMakeLen(b, 4)
}

func BenchmarkMakeLen_8(b *testing.B) {
	benchmarkMakeLen(b, 8)
}

func BenchmarkMakeLen_16(b *testing.B) {
	benchmarkMakeLen(b, 16)
}

func BenchmarkMakeLen_32(b *testing.B) {
	benchmarkMakeLen(b, 32)
}

func BenchmarkMakeLen_64(b *testing.B) {
	benchmarkMakeLen(b, 64)
}

func BenchmarkMakeLen_128(b *testing.B) {
	benchmarkMakeLen(b, 128)
}

func BenchmarkMakeLen_256(b *testing.B) {
	benchmarkMakeLen(b, 256)
}

func BenchmarkMakeLen_512(b *testing.B) {
	benchmarkMakeLen(b, 512)
}

func BenchmarkMakeLen_1024(b *testing.B) {
	benchmarkMakeLen(b, 1024)
}

func BenchmarkMakeLen_2048(b *testing.B) {
	benchmarkMakeLen(b, 2048)
}

func BenchmarkMakeLen_4096(b *testing.B) {
	benchmarkMakeLen(b, 4096)
}

func BenchmarkMakeLen_8192(b *testing.B) {
	benchmarkMakeLen(b, 8192)
}
