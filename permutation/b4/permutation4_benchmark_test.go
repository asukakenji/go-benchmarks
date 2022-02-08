package b4_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/permutation/b4"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkUintSliceGenerator(b *testing.B) {
	bcommon.Reset()
	benchmark.CalibrateUintSliceGenerator(b, bcommon.NextSlice)
}

// BenchmarkPermutation4Order0-8   	    3000	    437654 ns/op
// BenchmarkPermutation4Order1-8   	    3000	    443764 ns/op
// BenchmarkPermutation4Order2-8   	    3000	    431026 ns/op <- Best
// BenchmarkPermutation4Order3-8   	    3000	    434179 ns/op <- Best
// BenchmarkPermutation4Order4-8   	    3000	    442334 ns/op
// BenchmarkPermutation4Order5-8   	    3000	    430472 ns/op <- Best

func BenchmarkPermutation4Order0(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b4.Permutation4ParamOrder0[uint])
}

func BenchmarkPermutation4Order1(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b4.Permutation4ParamOrder1[uint])
}

func BenchmarkPermutation4Order2(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b4.Permutation4ParamOrder2[uint])
}

func BenchmarkPermutation4Order3(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b4.Permutation4ParamOrder3[uint])
}

func BenchmarkPermutation4Order4(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b4.Permutation4ParamOrder4[uint])
}

func BenchmarkPermutation4Order5(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b4.Permutation4ParamOrder5[uint])
}
