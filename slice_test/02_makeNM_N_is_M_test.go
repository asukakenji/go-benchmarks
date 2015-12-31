package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeNM_1_1(b *testing.B) {
	benchmarkMakeNM(b, 1, 1)
}

func BenchmarkMakeNM_2_2(b *testing.B) {
	benchmarkMakeNM(b, 2, 2)
}

func BenchmarkMakeNM_4_4(b *testing.B) {
	benchmarkMakeNM(b, 4, 4)
}

func BenchmarkMakeNM_8_8(b *testing.B) {
	benchmarkMakeNM(b, 8, 8)
}

func BenchmarkMakeNM_16_16(b *testing.B) {
	benchmarkMakeNM(b, 16, 16)
}

func BenchmarkMakeNM_32_32(b *testing.B) {
	benchmarkMakeNM(b, 32, 32)
}

func BenchmarkMakeNM_64_64(b *testing.B) {
	benchmarkMakeNM(b, 64, 64)
}

func BenchmarkMakeNM_128_128(b *testing.B) {
	benchmarkMakeNM(b, 128, 128)
}

func BenchmarkMakeNM_256_256(b *testing.B) {
	benchmarkMakeNM(b, 256, 256)
}

func BenchmarkMakeNM_512_512(b *testing.B) {
	benchmarkMakeNM(b, 512, 512)
}

func BenchmarkMakeNM_1024_1024(b *testing.B) {
	benchmarkMakeNM(b, 1024, 1024)
}

func BenchmarkMakeNM_2048_2048(b *testing.B) {
	benchmarkMakeNM(b, 2048, 2048)
}

func BenchmarkMakeNM_4096_4096(b *testing.B) {
	benchmarkMakeNM(b, 4096, 4096)
}

func BenchmarkMakeNM_8192_8192(b *testing.B) {
	benchmarkMakeNM(b, 8192, 8192)
}
