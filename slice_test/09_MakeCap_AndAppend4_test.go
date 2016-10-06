package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeCap_0_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 0)
}

func BenchmarkMakeCap_1_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 1)
}

func BenchmarkMakeCap_2_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 2)
}

func BenchmarkMakeCap_4_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 4)
}

func BenchmarkMakeCap_8_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 8)
}

func BenchmarkMakeCap_16_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 16)
}

func BenchmarkMakeCap_32_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 32)
}

func BenchmarkMakeCap_64_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 64)
}

func BenchmarkMakeCap_128_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 128)
}

func BenchmarkMakeCap_256_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 256)
}

func BenchmarkMakeCap_512_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 512)
}

func BenchmarkMakeCap_1024_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 1024)
}

func BenchmarkMakeCap_2048_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 2048)
}

func BenchmarkMakeCap_4096_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 4096)
}

func BenchmarkMakeCap_8192_AndAppend4(b *testing.B) {
	benchmarkMakeCapAndAppend4(b.N, 8192)
}
