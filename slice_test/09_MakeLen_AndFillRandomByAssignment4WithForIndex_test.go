package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

// Length not enough for assignment
// func BenchmarkMakeLen_0_0_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
// 	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 0)
// }

// Length not enough for assignment
// func BenchmarkMakeLen_1_1_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
// 	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 1)
// }

// Length not enough for assignment
// func BenchmarkMakeLen_2_2_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
// 	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 2)
// }

func BenchmarkMakeLen_4_4_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 4)
}

func BenchmarkMakeLen_8_8_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 8)
}

func BenchmarkMakeLen_16_16_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 16)
}

func BenchmarkMakeLen_32_32_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 32)
}

func BenchmarkMakeLen_64_64_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 64)
}

func BenchmarkMakeLen_128_128_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 128)
}

func BenchmarkMakeLen_256_256_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 256)
}

func BenchmarkMakeLen_512_512_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 512)
}

func BenchmarkMakeLen_1024_1024_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 1024)
}

func BenchmarkMakeLen_2048_2048_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 2048)
}

func BenchmarkMakeLen_4096_4096_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 4096)
}

func BenchmarkMakeLen_8192_8192_AndFillRandomByAssignment4WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(b.N, 8192)
}
