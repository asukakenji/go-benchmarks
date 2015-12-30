package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeNMAndAssignOneElementMTimes(b *testing.B, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < b.N; i++ {
		arr := make([]int, length, capacity)
		for j := 0; j < capacity; j++ {
			arr[j] = gen()
		}
	}
}

func BenchmarkMakeNM_0_0_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 0, 0)
}

func BenchmarkMakeNM_1_1_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 1, 1)
}

func BenchmarkMakeNM_2_2_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 2, 2)
}

func BenchmarkMakeNM_4_4_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 4, 4)
}

func BenchmarkMakeNM_8_8_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 8, 8)
}

func BenchmarkMakeNM_16_16_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 16, 16)
}

func BenchmarkMakeNM_32_32_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 32, 32)
}

func BenchmarkMakeNM_64_64_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 64, 64)
}

func BenchmarkMakeNM_128_128_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 128, 128)
}

func BenchmarkMakeNM_256_256_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 256, 256)
}

func BenchmarkMakeNM_512_512_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 512, 512)
}

func BenchmarkMakeNM_1024_1024_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 1024, 1024)
}

func BenchmarkMakeNM_2048_2048_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 2048, 2048)
}

func BenchmarkMakeNM_4096_4096_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 4096, 4096)
}

func BenchmarkMakeNM_8192_8192_AndAssignOneElementMTimes(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimes(b, 8192, 8192)
}
