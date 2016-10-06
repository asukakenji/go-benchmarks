package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeLen_0_0_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 0)
}

func BenchmarkMakeLen_1_1_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 1)
}

func BenchmarkMakeLen_2_2_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 2)
}

func BenchmarkMakeLen_4_4_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 4)
}

func BenchmarkMakeLen_8_8_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 8)
}

func BenchmarkMakeLen_16_16_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 16)
}

func BenchmarkMakeLen_32_32_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 32)
}

func BenchmarkMakeLen_64_64_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 64)
}

func BenchmarkMakeLen_128_128_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 128)
}

func BenchmarkMakeLen_256_256_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 256)
}

func BenchmarkMakeLen_512_512_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 512)
}

func BenchmarkMakeLen_1024_1024_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 1024)
}

func BenchmarkMakeLen_2048_2048_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 2048)
}

func BenchmarkMakeLen_4096_4096_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 4096)
}

func BenchmarkMakeLen_8192_8192_AndFillRandomByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(b.N, 8192)
}
