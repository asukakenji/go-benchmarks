package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeNMAndAppendOneElement(b *testing.B, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < b.N; i++ {
		arr := make([]int, length, capacity)
		arr = append(arr, gen())
	}
}

func BenchmarkMakeNM_0_1_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 1)
}

func BenchmarkMakeNM_0_2_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 2)
}

func BenchmarkMakeNM_0_4_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 4)
}

func BenchmarkMakeNM_0_8_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 8)
}

func BenchmarkMakeNM_0_16_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 16)
}

func BenchmarkMakeNM_0_32_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 32)
}

func BenchmarkMakeNM_0_64_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 64)
}

func BenchmarkMakeNM_0_128_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 128)
}

func BenchmarkMakeNM_0_256_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 256)
}

func BenchmarkMakeNM_0_512_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 512)
}

func BenchmarkMakeNM_0_1024_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 1024)
}

func BenchmarkMakeNM_0_2048_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 2048)
}

func BenchmarkMakeNM_0_4096_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 4096)
}

func BenchmarkMakeNM_0_8192_AndAppendOneElement(b *testing.B) {
	benchmarkMakeNMAndAppendOneElement(b, 0, 8192)
}
