package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeNMAndAppendOneElementMTimes(b *testing.B, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < b.N; i++ {
		arr := make([]int, length, capacity)
		for j := 0; j < capacity; j++ {
			arr = append(arr, gen())
		}
	}
}

func BenchmarkMakeNM_0_0_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 0)
}

func BenchmarkMakeNM_0_1_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 1)
}

func BenchmarkMakeNM_0_2_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 2)
}

func BenchmarkMakeNM_0_4_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 4)
}

func BenchmarkMakeNM_0_8_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 8)
}

func BenchmarkMakeNM_0_16_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 16)
}

func BenchmarkMakeNM_0_32_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 32)
}

func BenchmarkMakeNM_0_64_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 64)
}

func BenchmarkMakeNM_0_128_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 128)
}

func BenchmarkMakeNM_0_256_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 256)
}

func BenchmarkMakeNM_0_512_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 512)
}

func BenchmarkMakeNM_0_1024_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 1024)
}

func BenchmarkMakeNM_0_2048_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 2048)
}

func BenchmarkMakeNM_0_4096_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 4096)
}

func BenchmarkMakeNM_0_8192_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAppendOneElementMTimes(b, 0, 8192)
}
