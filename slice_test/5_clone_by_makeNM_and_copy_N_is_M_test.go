package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkCloneByMakeNM_1_1_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 1, 1)
}

func BenchmarkCloneByMakeNM_2_2_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 2, 2)
}

func BenchmarkCloneByMakeNM_4_4_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 4, 4)
}

func BenchmarkCloneByMakeNM_8_8_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 8, 8)
}

func BenchmarkCloneByMakeNM_16_16_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 16, 16)
}

func BenchmarkCloneByMakeNM_32_32_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 32, 32)
}

func BenchmarkCloneByMakeNM_64_64_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 64, 64)
}

func BenchmarkCloneByMakeNM_128_128_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 128, 128)
}

func BenchmarkCloneByMakeNM_256_256_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 256, 256)
}

func BenchmarkCloneByMakeNM_512_512_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 512, 512)
}

func BenchmarkCloneByMakeNM_1024_1024_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 1024, 1024)
}

func BenchmarkCloneByMakeNM_2048_2048_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 2048, 2048)
}

func BenchmarkCloneByMakeNM_4096_4096_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 4096, 4096)
}

func BenchmarkCloneByMakeNM_8192_8192_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b, 8192, 8192)
}
