package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeLen_0_0_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 0)
}

func BenchmarkMakeLen_1_1_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 1)
}

func BenchmarkMakeLen_2_2_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 2)
}

func BenchmarkMakeLen_4_4_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 4)
}

func BenchmarkMakeLen_8_8_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 8)
}

func BenchmarkMakeLen_16_16_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 16)
}

func BenchmarkMakeLen_32_32_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 32)
}

func BenchmarkMakeLen_64_64_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 64)
}

func BenchmarkMakeLen_128_128_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 128)
}

func BenchmarkMakeLen_256_256_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 256)
}

func BenchmarkMakeLen_512_512_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 512)
}

func BenchmarkMakeLen_1024_1024_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 1024)
}

func BenchmarkMakeLen_2048_2048_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 2048)
}

func BenchmarkMakeLen_4096_4096_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 4096)
}

func BenchmarkMakeLen_8192_8192_AndFillRandomByAssignmentWithForRange2(b *testing.B) {
	benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(b.N, 8192)
}
