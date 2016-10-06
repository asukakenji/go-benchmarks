package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeLenCapAndAppendTwoElements(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		arr = append(arr, gen(), gen())
	}
}

func BenchmarkMakeLenCap_0_2_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 2)
}

func BenchmarkMakeLenCap_0_4_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 4)
}

func BenchmarkMakeLenCap_0_8_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 8)
}

func BenchmarkMakeLenCap_0_16_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 16)
}

func BenchmarkMakeLenCap_0_32_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 32)
}

func BenchmarkMakeLenCap_0_64_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 64)
}

func BenchmarkMakeLenCap_0_128_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 128)
}

func BenchmarkMakeLenCap_0_256_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 256)
}

func BenchmarkMakeLenCap_0_512_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 512)
}

func BenchmarkMakeLenCap_0_1024_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 1024)
}

func BenchmarkMakeLenCap_0_2048_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 2048)
}

func BenchmarkMakeLenCap_0_4096_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 4096)
}

func BenchmarkMakeLenCap_0_8192_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeLenCapAndAppendTwoElements(b.N, 0, 8192)
}
