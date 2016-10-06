package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeLenCap_0_0_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 0)
}

func BenchmarkMakeLenCap_0_1_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 1)
}

func BenchmarkMakeLenCap_0_2_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 2)
}

func BenchmarkMakeLenCap_0_4_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 4)
}

func BenchmarkMakeLenCap_0_8_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 8)
}

func BenchmarkMakeLenCap_0_16_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 16)
}

func BenchmarkMakeLenCap_0_32_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 32)
}

func BenchmarkMakeLenCap_0_64_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 64)
}

func BenchmarkMakeLenCap_0_128_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 128)
}

func BenchmarkMakeLenCap_0_256_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 256)
}

func BenchmarkMakeLenCap_0_512_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 512)
}

func BenchmarkMakeLenCap_0_1024_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 1024)
}

func BenchmarkMakeLenCap_0_2048_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 2048)
}

func BenchmarkMakeLenCap_0_4096_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 4096)
}

func BenchmarkMakeLenCap_0_8192_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(b.N, 0, 8192)
}
