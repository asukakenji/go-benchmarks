package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeLenCap_0_0_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 0)
}

func BenchmarkMakeLenCap_0_1_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 1)
}

func BenchmarkMakeLenCap_0_2_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 2)
}

func BenchmarkMakeLenCap_0_4_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 4)
}

func BenchmarkMakeLenCap_0_8_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 8)
}

func BenchmarkMakeLenCap_0_16_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 16)
}

func BenchmarkMakeLenCap_0_32_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 32)
}

func BenchmarkMakeLenCap_0_64_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 64)
}

func BenchmarkMakeLenCap_0_128_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 128)
}

func BenchmarkMakeLenCap_0_256_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 256)
}

func BenchmarkMakeLenCap_0_512_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 512)
}

func BenchmarkMakeLenCap_0_1024_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 1024)
}

func BenchmarkMakeLenCap_0_2048_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 2048)
}

func BenchmarkMakeLenCap_0_4096_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 4096)
}

func BenchmarkMakeLenCap_0_8192_AndAppend4(b *testing.B) {
	benchmarkMakeLenCapAndAppend4(b.N, 0, 8192)
}
