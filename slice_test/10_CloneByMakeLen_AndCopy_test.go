package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkCloneByMakeLen_0_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 0)
}

func BenchmarkCloneByMakeLen_1_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 1)
}

func BenchmarkCloneByMakeLen_2_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 2)
}

func BenchmarkCloneByMakeLen_4_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 4)
}

func BenchmarkCloneByMakeLen_8_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 8)
}

func BenchmarkCloneByMakeLen_16_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 16)
}

func BenchmarkCloneByMakeLen_32_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 32)
}

func BenchmarkCloneByMakeLen_64_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 64)
}

func BenchmarkCloneByMakeLen_128_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 128)
}

func BenchmarkCloneByMakeLen_256_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 256)
}

func BenchmarkCloneByMakeLen_512_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 512)
}

func BenchmarkCloneByMakeLen_1024_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 1024)
}

func BenchmarkCloneByMakeLen_2048_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 2048)
}

func BenchmarkCloneByMakeLen_4096_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 4096)
}

func BenchmarkCloneByMakeLen_8192_AndCopy(b *testing.B) {
	benchmarkCloneByMakeLenAndCopy(b.N, 8192)
}
