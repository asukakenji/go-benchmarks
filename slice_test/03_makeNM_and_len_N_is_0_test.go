package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeLenCap_0_0_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 0)
}

func BenchmarkMakeLenCap_0_1_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 1)
}

func BenchmarkMakeLenCap_0_2_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 2)
}

func BenchmarkMakeLenCap_0_4_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 4)
}

func BenchmarkMakeLenCap_0_8_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 8)
}

func BenchmarkMakeLenCap_0_16_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 16)
}

func BenchmarkMakeLenCap_0_32_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 32)
}

func BenchmarkMakeLenCap_0_64_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 64)
}

func BenchmarkMakeLenCap_0_128_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 128)
}

func BenchmarkMakeLenCap_0_256_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 256)
}

func BenchmarkMakeLenCap_0_512_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 512)
}

func BenchmarkMakeLenCap_0_1024_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 1024)
}

func BenchmarkMakeLenCap_0_2048_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 2048)
}

func BenchmarkMakeLenCap_0_4096_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 4096)
}

func BenchmarkMakeLenCap_0_8192_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 0, 8192)
}
