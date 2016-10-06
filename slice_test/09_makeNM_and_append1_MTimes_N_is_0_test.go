package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeLenCapAndAppendOneElementMTimes(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		for j := 0; j < capacity; j++ {
			arr = append(arr, gen())
		}
	}
}

func BenchmarkMakeLenCap_0_0_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 0)
}

func BenchmarkMakeLenCap_0_1_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 1)
}

func BenchmarkMakeLenCap_0_2_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 2)
}

func BenchmarkMakeLenCap_0_4_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 4)
}

func BenchmarkMakeLenCap_0_8_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 8)
}

func BenchmarkMakeLenCap_0_16_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 16)
}

func BenchmarkMakeLenCap_0_32_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 32)
}

func BenchmarkMakeLenCap_0_64_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 64)
}

func BenchmarkMakeLenCap_0_128_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 128)
}

func BenchmarkMakeLenCap_0_256_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 256)
}

func BenchmarkMakeLenCap_0_512_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 512)
}

func BenchmarkMakeLenCap_0_1024_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 1024)
}

func BenchmarkMakeLenCap_0_2048_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 2048)
}

func BenchmarkMakeLenCap_0_4096_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 4096)
}

func BenchmarkMakeLenCap_0_8192_AndAppendOneElementMTimes(b *testing.B) {
	benchmarkMakeLenCapAndAppendOneElementMTimes(b.N, 0, 8192)
}
