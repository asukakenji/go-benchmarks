package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

// func BenchmarkMakeLenCap_0_0(b *testing.B) {
// 	benchmarkMakeLenCap(b.N, 0, 0)
// }

func BenchmarkMakeLenCap_1_1(b *testing.B) {
	benchmarkMakeLenCap(b.N, 1, 1)
}

func BenchmarkMakeLenCap_2_2(b *testing.B) {
	benchmarkMakeLenCap(b.N, 2, 2)
}

func BenchmarkMakeLenCap_4_4(b *testing.B) {
	benchmarkMakeLenCap(b.N, 4, 4)
}

func BenchmarkMakeLenCap_8_8(b *testing.B) {
	benchmarkMakeLenCap(b.N, 8, 8)
}

func BenchmarkMakeLenCap_16_16(b *testing.B) {
	benchmarkMakeLenCap(b.N, 16, 16)
}

func BenchmarkMakeLenCap_32_32(b *testing.B) {
	benchmarkMakeLenCap(b.N, 32, 32)
}

func BenchmarkMakeLenCap_64_64(b *testing.B) {
	benchmarkMakeLenCap(b.N, 64, 64)
}

func BenchmarkMakeLenCap_128_128(b *testing.B) {
	benchmarkMakeLenCap(b.N, 128, 128)
}

func BenchmarkMakeLenCap_256_256(b *testing.B) {
	benchmarkMakeLenCap(b.N, 256, 256)
}

func BenchmarkMakeLenCap_512_512(b *testing.B) {
	benchmarkMakeLenCap(b.N, 512, 512)
}

func BenchmarkMakeLenCap_1024_1024(b *testing.B) {
	benchmarkMakeLenCap(b.N, 1024, 1024)
}

func BenchmarkMakeLenCap_2048_2048(b *testing.B) {
	benchmarkMakeLenCap(b.N, 2048, 2048)
}

func BenchmarkMakeLenCap_4096_4096(b *testing.B) {
	benchmarkMakeLenCap(b.N, 4096, 4096)
}

func BenchmarkMakeLenCap_8192_8192(b *testing.B) {
	benchmarkMakeLenCap(b.N, 8192, 8192)
}
