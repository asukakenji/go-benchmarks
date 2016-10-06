package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkCloneByAppendMElementsToNilSlice_1(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 1)
}

func BenchmarkCloneByAppendMElementsToNilSlice_2(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 2)
}

func BenchmarkCloneByAppendMElementsToNilSlice_4(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 4)
}

func BenchmarkCloneByAppendMElementsToNilSlice_8(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 8)
}

func BenchmarkCloneByAppendMElementsToNilSlice_16(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 16)
}

func BenchmarkCloneByAppendMElementsToNilSlice_32(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 32)
}

func BenchmarkCloneByAppendMElementsToNilSlice_64(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 64)
}

func BenchmarkCloneByAppendMElementsToNilSlice_128(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 128)
}

func BenchmarkCloneByAppendMElementsToNilSlice_256(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 256)
}

func BenchmarkCloneByAppendMElementsToNilSlice_512(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 512)
}

func BenchmarkCloneByAppendMElementsToNilSlice_1024(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 1024)
}

func BenchmarkCloneByAppendMElementsToNilSlice_2048(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 2048)
}

func BenchmarkCloneByAppendMElementsToNilSlice_4096(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 4096)
}

func BenchmarkCloneByAppendMElementsToNilSlice_8192(b *testing.B) {
	benchmarkCloneByAppendMElementsToNilSlice(b.N, 8192)
}
