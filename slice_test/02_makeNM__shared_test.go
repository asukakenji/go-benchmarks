package slice_test

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func benchmarkMakeNM(b *testing.B, length, capacity int) {
	for i := 0; i < b.N; i++ {
		_ = make([]int, length, capacity)
	}
}
