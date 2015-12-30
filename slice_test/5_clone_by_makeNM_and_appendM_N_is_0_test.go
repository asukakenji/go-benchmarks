package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkCloneByMakeNMAndAppendMElements(b *testing.B, length, capacity int) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, length, capacity)
		arr = append(arr, numbers[0:capacity]...)
	}
}

func BenchmarkCloneByMakeNM_0_0_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 0)
}

func BenchmarkCloneByMakeNM_0_1_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 1)
}

func BenchmarkCloneByMakeNM_0_2_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 2)
}

func BenchmarkCloneByMakeNM_0_4_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 4)
}

func BenchmarkCloneByMakeNM_0_8_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 8)
}

func BenchmarkCloneByMakeNM_0_16_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 16)
}

func BenchmarkCloneByMakeNM_0_32_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 32)
}

func BenchmarkCloneByMakeNM_0_64_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 64)
}

func BenchmarkCloneByMakeNM_0_128_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 128)
}

func BenchmarkCloneByMakeNM_0_256_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 256)
}

func BenchmarkCloneByMakeNM_0_512_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 512)
}

func BenchmarkCloneByMakeNM_0_1024_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 1024)
}

func BenchmarkCloneByMakeNM_0_2048_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 2048)
}

func BenchmarkCloneByMakeNM_0_4096_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 4096)
}

func BenchmarkCloneByMakeNM_0_8192_AndAppendMElements(b *testing.B) {
	benchmarkCloneByMakeNMAndAppendMElements(b, 0, 8192)
}
