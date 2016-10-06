package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkCloneByMakeCap_0_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 0)
}

func BenchmarkCloneByMakeCap_1_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 1)
}

func BenchmarkCloneByMakeCap_2_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 2)
}

func BenchmarkCloneByMakeCap_4_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 4)
}

func BenchmarkCloneByMakeCap_8_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 8)
}

func BenchmarkCloneByMakeCap_16_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 16)
}

func BenchmarkCloneByMakeCap_32_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 32)
}

func BenchmarkCloneByMakeCap_64_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 64)
}

func BenchmarkCloneByMakeCap_128_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 128)
}

func BenchmarkCloneByMakeCap_256_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 256)
}

func BenchmarkCloneByMakeCap_512_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 512)
}

func BenchmarkCloneByMakeCap_1024_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 1024)
}

func BenchmarkCloneByMakeCap_2048_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 2048)
}

func BenchmarkCloneByMakeCap_4096_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 4096)
}

func BenchmarkCloneByMakeCap_8192_AndAppendN(b *testing.B) {
	benchmarkCloneByMakeCapAndAppendN(b.N, 8192)
}
