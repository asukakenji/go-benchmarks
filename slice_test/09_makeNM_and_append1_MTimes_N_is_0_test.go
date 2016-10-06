package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeLenCap_0_0_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 0)
}

func BenchmarkMakeLenCap_0_1_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 1)
}

func BenchmarkMakeLenCap_0_2_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 2)
}

func BenchmarkMakeLenCap_0_4_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 4)
}

func BenchmarkMakeLenCap_0_8_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 8)
}

func BenchmarkMakeLenCap_0_16_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 16)
}

func BenchmarkMakeLenCap_0_32_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 32)
}

func BenchmarkMakeLenCap_0_64_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 64)
}

func BenchmarkMakeLenCap_0_128_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 128)
}

func BenchmarkMakeLenCap_0_256_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 256)
}

func BenchmarkMakeLenCap_0_512_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 512)
}

func BenchmarkMakeLenCap_0_1024_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 1024)
}

func BenchmarkMakeLenCap_0_2048_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 2048)
}

func BenchmarkMakeLenCap_0_4096_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 4096)
}

func BenchmarkMakeLenCap_0_8192_AndFillByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillByAppend1WithForIndex(b.N, 0, 8192)
}
