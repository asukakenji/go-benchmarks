package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeNM_0_0_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 0)
}

func BenchmarkMakeNM_0_1_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 1)
}

func BenchmarkMakeNM_0_2_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 2)
}

func BenchmarkMakeNM_0_4_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 4)
}

func BenchmarkMakeNM_0_8_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 8)
}

func BenchmarkMakeNM_0_16_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 16)
}

func BenchmarkMakeNM_0_32_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 32)
}

func BenchmarkMakeNM_0_64_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 64)
}

func BenchmarkMakeNM_0_128_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 128)
}

func BenchmarkMakeNM_0_256_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 256)
}

func BenchmarkMakeNM_0_512_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 512)
}

func BenchmarkMakeNM_0_1024_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 1024)
}

func BenchmarkMakeNM_0_2048_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 2048)
}

func BenchmarkMakeNM_0_4096_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 4096)
}

func BenchmarkMakeNM_0_8192_AndLen(b *testing.B) {
	benchmarkMakeNMAndLen(b, 0, 8192)
}
