package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeNMAndAssignOneElementMTimesWithRange1(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		for j := range arr {
			arr[j] = gen()
		}
	}
}

func BenchmarkMakeNM_0_0_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 0, 0)
}

func BenchmarkMakeNM_1_1_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 1, 1)
}

func BenchmarkMakeNM_2_2_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 2, 2)
}

func BenchmarkMakeNM_4_4_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 4, 4)
}

func BenchmarkMakeNM_8_8_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 8, 8)
}

func BenchmarkMakeNM_16_16_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 16, 16)
}

func BenchmarkMakeNM_32_32_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 32, 32)
}

func BenchmarkMakeNM_64_64_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 64, 64)
}

func BenchmarkMakeNM_128_128_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 128, 128)
}

func BenchmarkMakeNM_256_256_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 256, 256)
}

func BenchmarkMakeNM_512_512_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 512, 512)
}

func BenchmarkMakeNM_1024_1024_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 1024, 1024)
}

func BenchmarkMakeNM_2048_2048_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 2048, 2048)
}

func BenchmarkMakeNM_4096_4096_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 4096, 4096)
}

func BenchmarkMakeNM_8192_8192_AndAssignOneElementMTimesWithRange1(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange1(b.N, 8192, 8192)
}
