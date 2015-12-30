package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeNM_1_1_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 1, 1)
}

func BenchmarkMakeNM_2_2_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 2, 2)
}

func BenchmarkMakeNM_4_4_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 4, 4)
}

func BenchmarkMakeNM_8_8_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 8, 8)
}

func BenchmarkMakeNM_16_16_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 16, 16)
}

func BenchmarkMakeNM_32_32_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 32, 32)
}

func BenchmarkMakeNM_64_64_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 64, 64)
}

func BenchmarkMakeNM_128_128_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 128, 128)
}

func BenchmarkMakeNM_256_256_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 256, 256)
}

func BenchmarkMakeNM_512_512_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 512, 512)
}

func BenchmarkMakeNM_1024_1024_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 1024, 1024)
}

func BenchmarkMakeNM_2048_2048_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 2048, 2048)
}

func BenchmarkMakeNM_4096_4096_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 4096, 4096)
}

func BenchmarkMakeNM_8192_8192_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 8192, 8192)
}
