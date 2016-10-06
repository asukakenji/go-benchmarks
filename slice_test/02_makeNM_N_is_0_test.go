package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeLenCap_0_0(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 0)
}

func BenchmarkMakeLenCap_0_1(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 1)
}

func BenchmarkMakeLenCap_0_2(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 2)
}

func BenchmarkMakeLenCap_0_4(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 4)
}

func BenchmarkMakeLenCap_0_8(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 8)
}

func BenchmarkMakeLenCap_0_16(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 16)
}

func BenchmarkMakeLenCap_0_32(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 32)
}

func BenchmarkMakeLenCap_0_64(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 64)
}

func BenchmarkMakeLenCap_0_128(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 128)
}

func BenchmarkMakeLenCap_0_256(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 256)
}

func BenchmarkMakeLenCap_0_512(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 512)
}

func BenchmarkMakeLenCap_0_1024(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 1024)
}

func BenchmarkMakeLenCap_0_2048(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 1024)
}

func BenchmarkMakeLenCap_0_4096(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 1024)
}

func BenchmarkMakeLenCap_0_8192(b *testing.B) {
	benchmarkMakeLenCap(b.N, 0, 1024)
}
