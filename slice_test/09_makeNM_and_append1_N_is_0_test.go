package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeLenCapAndAppendOneElement(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		arr = append(arr, gen())
	}
}

func BenchmarkMakeLenCap_0_1_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 1)
}

func BenchmarkMakeLenCap_0_2_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 2)
}

func BenchmarkMakeLenCap_0_4_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 4)
}

func BenchmarkMakeLenCap_0_8_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 8)
}

func BenchmarkMakeLenCap_0_16_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 16)
}

func BenchmarkMakeLenCap_0_32_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 32)
}

func BenchmarkMakeLenCap_0_64_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 64)
}

func BenchmarkMakeLenCap_0_128_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 128)
}

func BenchmarkMakeLenCap_0_256_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 256)
}

func BenchmarkMakeLenCap_0_512_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 512)
}

func BenchmarkMakeLenCap_0_1024_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 1024)
}

func BenchmarkMakeLenCap_0_2048_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 2048)
}

func BenchmarkMakeLenCap_0_4096_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 4096)
}

func BenchmarkMakeLenCap_0_8192_AndAppendOneElement(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElement(b.N, 0, 8192)
}
