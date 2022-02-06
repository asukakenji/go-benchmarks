package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkCap_0_0(b *testing.B) {
	benchmarkCap(b, 0, 0)
}

func BenchmarkCap_0_1(b *testing.B) {
	benchmarkCap(b, 0, 1)
}

func BenchmarkCap_0_2(b *testing.B) {
	benchmarkCap(b, 0, 2)
}

func BenchmarkCap_0_4(b *testing.B) {
	benchmarkCap(b, 0, 4)
}

func BenchmarkCap_0_8(b *testing.B) {
	benchmarkCap(b, 0, 8)
}

func BenchmarkCap_0_16(b *testing.B) {
	benchmarkCap(b, 0, 16)
}

func BenchmarkCap_0_32(b *testing.B) {
	benchmarkCap(b, 0, 32)
}

func BenchmarkCap_0_64(b *testing.B) {
	benchmarkCap(b, 0, 64)
}

func BenchmarkCap_0_128(b *testing.B) {
	benchmarkCap(b, 0, 128)
}

func BenchmarkCap_0_256(b *testing.B) {
	benchmarkCap(b, 0, 256)
}

func BenchmarkCap_0_512(b *testing.B) {
	benchmarkCap(b, 0, 512)
}

func BenchmarkCap_0_1024(b *testing.B) {
	benchmarkCap(b, 0, 1024)
}

func BenchmarkCap_0_2048(b *testing.B) {
	benchmarkCap(b, 0, 2048)
}

func BenchmarkCap_0_4096(b *testing.B) {
	benchmarkCap(b, 0, 4096)
}

func BenchmarkCap_0_8192(b *testing.B) {
	benchmarkCap(b, 0, 8192)
}
