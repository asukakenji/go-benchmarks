package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func BenchmarkCloneByAppendNToNil_0(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 0)
}

func BenchmarkCloneByAppendNToNil_1(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 1)
}

func BenchmarkCloneByAppendNToNil_2(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 2)
}

func BenchmarkCloneByAppendNToNil_4(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 4)
}

func BenchmarkCloneByAppendNToNil_8(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 8)
}

func BenchmarkCloneByAppendNToNil_16(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 16)
}

func BenchmarkCloneByAppendNToNil_32(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 32)
}

func BenchmarkCloneByAppendNToNil_64(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 64)
}

func BenchmarkCloneByAppendNToNil_128(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 128)
}

func BenchmarkCloneByAppendNToNil_256(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 256)
}

func BenchmarkCloneByAppendNToNil_512(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 512)
}

func BenchmarkCloneByAppendNToNil_1024(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 1024)
}

func BenchmarkCloneByAppendNToNil_2048(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 2048)
}

func BenchmarkCloneByAppendNToNil_4096(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 4096)
}

func BenchmarkCloneByAppendNToNil_8192(b *testing.B) {
	benchmarkCloneByAppendNToNil(b.N, 8192)
}
