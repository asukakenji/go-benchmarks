package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeLenCap_1_1_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 1, 1)
}

func BenchmarkMakeLenCap_2_2_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 2, 2)
}

func BenchmarkMakeLenCap_4_4_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 4, 4)
}

func BenchmarkMakeLenCap_8_8_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 8, 8)
}

func BenchmarkMakeLenCap_16_16_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 16, 16)
}

func BenchmarkMakeLenCap_32_32_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 32, 32)
}

func BenchmarkMakeLenCap_64_64_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 64, 64)
}

func BenchmarkMakeLenCap_128_128_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 128, 128)
}

func BenchmarkMakeLenCap_256_256_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 256, 256)
}

func BenchmarkMakeLenCap_512_512_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 512, 512)
}

func BenchmarkMakeLenCap_1024_1024_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 1024, 1024)
}

func BenchmarkMakeLenCap_2048_2048_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 2048, 2048)
}

func BenchmarkMakeLenCap_4096_4096_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 4096, 4096)
}

func BenchmarkMakeLenCap_8192_8192_AndCap(b *testing.B) {
	benchmarkMakeLenCapAndCap(b.N, 8192, 8192)
}
