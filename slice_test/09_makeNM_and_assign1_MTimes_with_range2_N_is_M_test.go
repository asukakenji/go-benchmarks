package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeNMAndAssignOneElementMTimesWithRange2(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		for j, _ := range arr {
			arr[j] = gen()
		}
	}
}

func BenchmarkMakeNM_0_0_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 0, 0)
}

func BenchmarkMakeNM_1_1_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 1, 1)
}

func BenchmarkMakeNM_2_2_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 2, 2)
}

func BenchmarkMakeNM_4_4_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 4, 4)
}

func BenchmarkMakeNM_8_8_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 8, 8)
}

func BenchmarkMakeNM_16_16_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 16, 16)
}

func BenchmarkMakeNM_32_32_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 32, 32)
}

func BenchmarkMakeNM_64_64_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 64, 64)
}

func BenchmarkMakeNM_128_128_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 128, 128)
}

func BenchmarkMakeNM_256_256_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 256, 256)
}

func BenchmarkMakeNM_512_512_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 512, 512)
}

func BenchmarkMakeNM_1024_1024_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 1024, 1024)
}

func BenchmarkMakeNM_2048_2048_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 2048, 2048)
}

func BenchmarkMakeNM_4096_4096_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 4096, 4096)
}

func BenchmarkMakeNM_8192_8192_AndAssignOneElementMTimesWithRange2(b *testing.B) {
	benchmarkMakeNMAndAssignOneElementMTimesWithRange2(b.N, 8192, 8192)
}
