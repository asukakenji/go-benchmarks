package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

// Length not enough for assignment
// func BenchmarkMakeLen_0_0_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
// 	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 0)
// }

// Length not enough for assignment
// func BenchmarkMakeLen_1_1_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
// 	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 1)
// }

func BenchmarkMakeLen_2_2_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 2)
}

func BenchmarkMakeLen_4_4_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 4)
}

func BenchmarkMakeLen_8_8_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 8)
}

func BenchmarkMakeLen_16_16_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 16)
}

func BenchmarkMakeLen_32_32_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 32)
}

func BenchmarkMakeLen_64_64_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 64)
}

func BenchmarkMakeLen_128_128_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 128)
}

func BenchmarkMakeLen_256_256_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 256)
}

func BenchmarkMakeLen_512_512_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 512)
}

func BenchmarkMakeLen_1024_1024_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 1024)
}

func BenchmarkMakeLen_2048_2048_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 2048)
}

func BenchmarkMakeLen_4096_4096_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 4096)
}

func BenchmarkMakeLen_8192_8192_AndFillRandomByAssignment2WithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(b.N, 8192)
}
