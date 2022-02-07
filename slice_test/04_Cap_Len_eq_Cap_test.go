package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

// func BenchmarkCap_0_0(b *testing.B) {
// 	benchmarkCap(b.N, numbers[0:0:0])
// }

func BenchmarkCap_1_1(b *testing.B) {
	benchmarkCap(b.N, numbers[0:1:1])
}

func BenchmarkCap_2_2(b *testing.B) {
	benchmarkCap(b.N, numbers[0:2:2])
}

func BenchmarkCap_4_4(b *testing.B) {
	benchmarkCap(b.N, numbers[0:4:4])
}

func BenchmarkCap_8_8(b *testing.B) {
	benchmarkCap(b.N, numbers[0:8:8])
}

func BenchmarkCap_16_16(b *testing.B) {
	benchmarkCap(b.N, numbers[0:16:16])
}

func BenchmarkCap_32_32(b *testing.B) {
	benchmarkCap(b.N, numbers[0:32:32])
}

func BenchmarkCap_64_64(b *testing.B) {
	benchmarkCap(b.N, numbers[0:64:64])
}

func BenchmarkCap_128_128(b *testing.B) {
	benchmarkCap(b.N, numbers[0:128:128])
}

func BenchmarkCap_256_256(b *testing.B) {
	benchmarkCap(b.N, numbers[0:256:256])
}

func BenchmarkCap_512_512(b *testing.B) {
	benchmarkCap(b.N, numbers[0:512:512])
}

func BenchmarkCap_1024_1024(b *testing.B) {
	benchmarkCap(b.N, numbers[0:1024:1024])
}

func BenchmarkCap_2048_2048(b *testing.B) {
	benchmarkCap(b.N, numbers[0:2048:2048])
}

func BenchmarkCap_4096_4096(b *testing.B) {
	benchmarkCap(b.N, numbers[0:4096:4096])
}

func BenchmarkCap_8192_8192(b *testing.B) {
	benchmarkCap(b.N, numbers[0:8192:8192])
}
