package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

// func BenchmarkMakeLenCap_0_0_AndLen(b *testing.B) {
// 	benchmarkMakeLenCapAndLen(b.N, 0, 0)
// }

func BenchmarkMakeLenCap_1_1_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 1, 1)
}

func BenchmarkMakeLenCap_2_2_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 2, 2)
}

func BenchmarkMakeLenCap_4_4_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 4, 4)
}

func BenchmarkMakeLenCap_8_8_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 8, 8)
}

func BenchmarkMakeLenCap_16_16_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 16, 16)
}

func BenchmarkMakeLenCap_32_32_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 32, 32)
}

func BenchmarkMakeLenCap_64_64_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 64, 64)
}

func BenchmarkMakeLenCap_128_128_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 128, 128)
}

func BenchmarkMakeLenCap_256_256_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 256, 256)
}

func BenchmarkMakeLenCap_512_512_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 512, 512)
}

func BenchmarkMakeLenCap_1024_1024_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 1024, 1024)
}

func BenchmarkMakeLenCap_2048_2048_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 2048, 2048)
}

func BenchmarkMakeLenCap_4096_4096_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 4096, 4096)
}

func BenchmarkMakeLenCap_8192_8192_AndLen(b *testing.B) {
	benchmarkMakeLenCapAndLen(b.N, 8192, 8192)
}
