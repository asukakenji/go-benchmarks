package b3_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/permutation/b3"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkUintSliceGenerator(b *testing.B) {
	bcommon.Reset()
	benchmark.CalibrateUintSliceGenerator(b, bcommon.NextSlice)
}

// BenchmarkPermutation3Order0-8   	    3000	    511230 ns/op
// BenchmarkPermutation3Order1-8   	    3000	    510388 ns/op
// BenchmarkPermutation3Order2-8   	    3000	    511744 ns/op
// BenchmarkPermutation3Order3-8   	    3000	    505027 ns/op <- Best
// BenchmarkPermutation3Order4-8   	    3000	    518075 ns/op
// BenchmarkPermutation3Order5-8   	    3000	    500379 ns/op <- Best

func BenchmarkPermutation3Order0(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b3.Permutation3ParamOrder0[uint])
}

func BenchmarkPermutation3Order1(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b3.Permutation3ParamOrder1[uint])
}

func BenchmarkPermutation3Order2(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b3.Permutation3ParamOrder2[uint])
}

func BenchmarkPermutation3Order3(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b3.Permutation3ParamOrder3[uint])
}

func BenchmarkPermutation3Order4(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b3.Permutation3ParamOrder4[uint])
}

func BenchmarkPermutation3Order5(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b3.Permutation3ParamOrder5[uint])
}
