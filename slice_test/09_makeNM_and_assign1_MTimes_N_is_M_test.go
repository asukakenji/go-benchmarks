package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeLen_0_0_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 0)
}

func BenchmarkMakeLen_1_1_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 1)
}

func BenchmarkMakeLen_2_2_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 2)
}

func BenchmarkMakeLen_4_4_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 4)
}

func BenchmarkMakeLen_8_8_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 8)
}

func BenchmarkMakeLen_16_16_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 16)
}

func BenchmarkMakeLen_32_32_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 32)
}

func BenchmarkMakeLen_64_64_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 64)
}

func BenchmarkMakeLen_128_128_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 128)
}

func BenchmarkMakeLen_256_256_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 256)
}

func BenchmarkMakeLen_512_512_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 512)
}

func BenchmarkMakeLen_1024_1024_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 1024)
}

func BenchmarkMakeLen_2048_2048_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 2048)
}

func BenchmarkMakeLen_4096_4096_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 4096)
}

func BenchmarkMakeLen_8192_8192_AndFillByAssignmentWithForIndex(b *testing.B) {
	benchmarkMakeLenAndFillByAssignmentWithForIndex(b.N, 8192)
}
