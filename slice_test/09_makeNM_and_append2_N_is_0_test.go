package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeLenCap_0_2_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 2)
}

func BenchmarkMakeLenCap_0_4_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 4)
}

func BenchmarkMakeLenCap_0_8_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 8)
}

func BenchmarkMakeLenCap_0_16_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 16)
}

func BenchmarkMakeLenCap_0_32_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 32)
}

func BenchmarkMakeLenCap_0_64_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 64)
}

func BenchmarkMakeLenCap_0_128_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 128)
}

func BenchmarkMakeLenCap_0_256_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 256)
}

func BenchmarkMakeLenCap_0_512_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 512)
}

func BenchmarkMakeLenCap_0_1024_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 1024)
}

func BenchmarkMakeLenCap_0_2048_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 2048)
}

func BenchmarkMakeLenCap_0_4096_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 4096)
}

func BenchmarkMakeLenCap_0_8192_AndAppend2(b *testing.B) {
	benchmarkMakeLenCapAndAppend2(b.N, 0, 8192)
}
