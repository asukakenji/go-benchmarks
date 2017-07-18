package b5_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/permutation/b5"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkUintSliceGenerator(b *testing.B) {
	bcommon.Reset()
	benchmark.CalibrateUintSliceGenerator(b, bcommon.NextSlice)
}

// BenchmarkPermutation4Order0-8   	    3000	    440636 ns/op
// BenchmarkPermutation4Order1-8   	    3000	    444581 ns/op
// BenchmarkPermutation4Order2-8   	    3000	    437564 ns/op <- Best
// BenchmarkPermutation4Order3-8   	    3000	    433765 ns/op <- Best
// BenchmarkPermutation4Order4-8   	    3000	    442313 ns/op
// BenchmarkPermutation4Order5-8   	    3000	    432704 ns/op <- Best

func BenchmarkPermutation5Inc(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b5.Permutation5Inc)
}
