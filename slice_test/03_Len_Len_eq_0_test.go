package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkLen_0_0(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:0])
}

func BenchmarkLen_0_1(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:1])
}

func BenchmarkLen_0_2(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:2])
}

func BenchmarkLen_0_4(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:4])
}

func BenchmarkLen_0_8(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:8])
}

func BenchmarkLen_0_16(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:16])
}

func BenchmarkLen_0_32(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:32])
}

func BenchmarkLen_0_64(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:64])
}

func BenchmarkLen_0_128(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:128])
}

func BenchmarkLen_0_256(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:256])
}

func BenchmarkLen_0_512(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:512])
}

func BenchmarkLen_0_1024(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:1024])
}

func BenchmarkLen_0_2048(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:2048])
}

func BenchmarkLen_0_4096(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:4096])
}

func BenchmarkLen_0_8192(b *testing.B) {
	benchmarkLen(b.N, numbers[0:0:8192])
}
