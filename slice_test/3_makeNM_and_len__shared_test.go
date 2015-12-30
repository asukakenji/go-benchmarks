package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeNMAndLen(b *testing.B, length, capacity int) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, length, capacity)
		_ = len(arr)
	}
}
