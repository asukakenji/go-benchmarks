package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkCloneByMakeNM_0_0_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 0)
}

func BenchmarkCloneByMakeNM_0_1_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 1)
}

func BenchmarkCloneByMakeNM_0_2_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 2)
}

func BenchmarkCloneByMakeNM_0_4_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 4)
}

func BenchmarkCloneByMakeNM_0_8_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 8)
}

func BenchmarkCloneByMakeNM_0_16_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 16)
}

func BenchmarkCloneByMakeNM_0_32_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 32)
}

func BenchmarkCloneByMakeNM_0_64_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 64)
}

func BenchmarkCloneByMakeNM_0_128_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 128)
}

func BenchmarkCloneByMakeNM_0_256_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 256)
}

func BenchmarkCloneByMakeNM_0_512_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 512)
}

func BenchmarkCloneByMakeNM_0_1024_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 1024)
}

func BenchmarkCloneByMakeNM_0_2048_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 2048)
}

func BenchmarkCloneByMakeNM_0_4096_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 4096)
}

func BenchmarkCloneByMakeNM_0_8192_AndCopy(b *testing.B) {
	benchmarkCloneByMakeNMAndCopy(b.N, 0, 8192)
}
