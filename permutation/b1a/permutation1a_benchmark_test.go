package b1a_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/permutation/b1a"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkUintSliceGenerator(b *testing.B) {
	bcommon.Reset()
	benchmark.CalibrateUintSliceGenerator(b, bcommon.NextSlice)
}

// BenchmarkPermutation1Order0-8   	    3000	    443578 ns/op
// BenchmarkPermutation1Order1-8   	    3000	    446210 ns/op
// BenchmarkPermutation1Order2-8   	    3000	    445479 ns/op
// BenchmarkPermutation1Order3-8   	    3000	    437096 ns/op <- Best
// BenchmarkPermutation1Order4-8   	    3000	    444762 ns/op
// BenchmarkPermutation1Order5-8   	    3000	    438581 ns/op <- Best

func BenchmarkPermutation1Order0(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1a.Permutation1ParamOrder0)
}

func BenchmarkPermutation1Order1(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1a.Permutation1ParamOrder1)
}

func BenchmarkPermutation1Order2(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1a.Permutation1ParamOrder2)
}

func BenchmarkPermutation1Order3(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1a.Permutation1ParamOrder3)
}

func BenchmarkPermutation1Order4(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1a.Permutation1ParamOrder4)
}

func BenchmarkPermutation1Order5(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1a.Permutation1ParamOrder5)
}
