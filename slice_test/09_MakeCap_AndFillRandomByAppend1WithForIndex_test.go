package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkMakeCap_0_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 0)
}

func BenchmarkMakeCap_1_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 1)
}

func BenchmarkMakeCap_2_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 2)
}

func BenchmarkMakeCap_4_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 4)
}

func BenchmarkMakeCap_8_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 8)
}

func BenchmarkMakeCap_16_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 16)
}

func BenchmarkMakeCap_32_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 32)
}

func BenchmarkMakeCap_64_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 64)
}

func BenchmarkMakeCap_128_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 128)
}

func BenchmarkMakeCap_256_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 256)
}

func BenchmarkMakeCap_512_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 512)
}

func BenchmarkMakeCap_1024_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 1024)
}

func BenchmarkMakeCap_2048_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 2048)
}

func BenchmarkMakeCap_4096_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 4096)
}

func BenchmarkMakeCap_8192_AndFillRandomByAppend1WithForIndex(b *testing.B) {
	benchmarkMakeCapAndFillRandomByAppend1WithForIndex(b.N, 8192)
}
