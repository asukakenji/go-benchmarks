package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeNMAndAppendTwoElements(b *testing.B, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < b.N; i++ {
		arr := make([]int, length, capacity)
		arr = append(arr, gen(), gen())
	}
}

func BenchmarkMakeNM_0_2_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 2)
}

func BenchmarkMakeNM_0_4_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 4)
}

func BenchmarkMakeNM_0_8_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 8)
}

func BenchmarkMakeNM_0_16_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 16)
}

func BenchmarkMakeNM_0_32_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 32)
}

func BenchmarkMakeNM_0_64_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 64)
}

func BenchmarkMakeNM_0_128_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 128)
}

func BenchmarkMakeNM_0_256_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 256)
}

func BenchmarkMakeNM_0_512_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 512)
}

func BenchmarkMakeNM_0_1024_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 1024)
}

func BenchmarkMakeNM_0_2048_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 2048)
}

func BenchmarkMakeNM_0_4096_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 4096)
}

func BenchmarkMakeNM_0_8192_AndAppendTwoElements(b *testing.B) {
	benchmarkMakeNMAndAppendTwoElements(b, 0, 8192)
}
