package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeNM_0_0(b *testing.B) {
	benchmarkMakeNM(b, 0, 0)
}

func BenchmarkMakeNM_0_1(b *testing.B) {
	benchmarkMakeNM(b, 0, 1)
}

func BenchmarkMakeNM_0_2(b *testing.B) {
	benchmarkMakeNM(b, 0, 2)
}

func BenchmarkMakeNM_0_4(b *testing.B) {
	benchmarkMakeNM(b, 0, 4)
}

func BenchmarkMakeNM_0_8(b *testing.B) {
	benchmarkMakeNM(b, 0, 8)
}

func BenchmarkMakeNM_0_16(b *testing.B) {
	benchmarkMakeNM(b, 0, 16)
}

func BenchmarkMakeNM_0_32(b *testing.B) {
	benchmarkMakeNM(b, 0, 32)
}

func BenchmarkMakeNM_0_64(b *testing.B) {
	benchmarkMakeNM(b, 0, 64)
}

func BenchmarkMakeNM_0_128(b *testing.B) {
	benchmarkMakeNM(b, 0, 128)
}

func BenchmarkMakeNM_0_256(b *testing.B) {
	benchmarkMakeNM(b, 0, 256)
}

func BenchmarkMakeNM_0_512(b *testing.B) {
	benchmarkMakeNM(b, 0, 512)
}

func BenchmarkMakeNM_0_1024(b *testing.B) {
	benchmarkMakeNM(b, 0, 1024)
}

func BenchmarkMakeNM_0_2048(b *testing.B) {
	benchmarkMakeNM(b, 0, 1024)
}

func BenchmarkMakeNM_0_4096(b *testing.B) {
	benchmarkMakeNM(b, 0, 1024)
}

func BenchmarkMakeNM_0_8192(b *testing.B) {
	benchmarkMakeNM(b, 0, 1024)
}
