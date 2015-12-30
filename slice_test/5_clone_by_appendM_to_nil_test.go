package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkCloneByAppendMElementsToNilSlice(b *testing.B, capacity int) {
	for i := 0; i < b.N; i++ {
		_ = append([]int(nil), numbers[0:capacity]...)
	}
}

func BenchmarkCloneByAppendMElementsToNilSlice_1(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 1)
}

func BenchmarkCloneByAppendMElementsToNilSlice_2(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 2)
}

func BenchmarkCloneByAppendMElementsToNilSlice_4(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 4)
}

func BenchmarkCloneByAppendMElementsToNilSlice_8(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 8)
}

func BenchmarkCloneByAppendMElementsToNilSlice_16(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 16)
}

func BenchmarkCloneByAppendMElementsToNilSlice_32(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 32)
}

func BenchmarkCloneByAppendMElementsToNilSlice_64(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 64)
}

func BenchmarkCloneByAppendMElementsToNilSlice_128(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 128)
}

func BenchmarkCloneByAppendMElementsToNilSlice_256(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 256)
}

func BenchmarkCloneByAppendMElementsToNilSlice_512(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 512)
}

func BenchmarkCloneByAppendMElementsToNilSlice_1024(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 1024)
}

func BenchmarkCloneByAppendMElementsToNilSlice_2048(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 2048)
}

func BenchmarkCloneByAppendMElementsToNilSlice_4096(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 4096)
}

func BenchmarkCloneByAppendMElementsToNilSlice_8192(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b, 8192)
}
