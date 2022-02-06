package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

// func BenchmarkLen_0_0(b *testing.B) {
// 	benchmarkMakeLenCapAndLen(b, 0, 0)
// }

func BenchmarkLen_1_1(b *testing.B) {
	benchmarkLen(b, 1, 1)
}

func BenchmarkLen_2_2(b *testing.B) {
	benchmarkLen(b, 2, 2)
}

func BenchmarkLen_4_4(b *testing.B) {
	benchmarkLen(b, 4, 4)
}

func BenchmarkLen_8_8(b *testing.B) {
	benchmarkLen(b, 8, 8)
}

func BenchmarkLen_16_16(b *testing.B) {
	benchmarkLen(b, 16, 16)
}

func BenchmarkLen_32_32(b *testing.B) {
	benchmarkLen(b, 32, 32)
}

func BenchmarkLen_64_64(b *testing.B) {
	benchmarkLen(b, 64, 64)
}

func BenchmarkLen_128_128(b *testing.B) {
	benchmarkLen(b, 128, 128)
}

func BenchmarkLen_256_256(b *testing.B) {
	benchmarkLen(b, 256, 256)
}

func BenchmarkLen_512_512(b *testing.B) {
	benchmarkLen(b, 512, 512)
}

func BenchmarkLen_1024_1024(b *testing.B) {
	benchmarkLen(b, 1024, 1024)
}

func BenchmarkLen_2048_2048(b *testing.B) {
	benchmarkLen(b, 2048, 2048)
}

func BenchmarkLen_4096_4096(b *testing.B) {
	benchmarkLen(b, 4096, 4096)
}

func BenchmarkLen_8192_8192(b *testing.B) {
	benchmarkLen(b, 8192, 8192)
}
