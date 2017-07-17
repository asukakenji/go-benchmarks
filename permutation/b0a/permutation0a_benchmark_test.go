package b0a_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/permutation/b0a"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkUintSliceGenerator(b *testing.B) {
	bcommon.Reset()
	benchmark.CalibrateUintSliceGenerator(b, bcommon.NextSlice)
}

// BenchmarkPermutation0Order0-8   	    3000	    489643 ns/op
// BenchmarkPermutation0Order1-8   	    3000	    492871 ns/op
// BenchmarkPermutation0Order2-8   	    3000	    482690 ns/op <- Best
// BenchmarkPermutation0Order3-8   	    3000	    485489 ns/op <- Best
// BenchmarkPermutation0Order4-8   	    3000	    500723 ns/op
// BenchmarkPermutation0Order5-8   	    3000	    482149 ns/op <- Best

func BenchmarkPermutation0Order0(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b0a.Permutation0ParamOrder0)
}

func BenchmarkPermutation0Order1(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b0a.Permutation0ParamOrder1)
}

func BenchmarkPermutation0Order2(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b0a.Permutation0ParamOrder2)
}

func BenchmarkPermutation0Order3(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b0a.Permutation0ParamOrder3)
}

func BenchmarkPermutation0Order4(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b0a.Permutation0ParamOrder4)
}

func BenchmarkPermutation0Order5(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b0a.Permutation0ParamOrder5)
}
