package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkCloneByMakeNMAndCopy(b *testing.B, length, capacity int) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, length, capacity)
		copy(arr, numbers[0:capacity])
	}
}
