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

// BenchmarkPermutation3Order0-8   	    3000	    440636 ns/op
// BenchmarkPermutation3Order1-8   	    3000	    444581 ns/op
// BenchmarkPermutation3Order2-8   	    3000	    437564 ns/op <- Best
// BenchmarkPermutation3Order3-8   	    3000	    433765 ns/op <- Best
// BenchmarkPermutation3Order4-8   	    3000	    442313 ns/op
// BenchmarkPermutation3Order5-8   	    3000	    432704 ns/op <- Best

func BenchmarkPermutation3Order0(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b3.Permutation3ParamOrder0)
}

func BenchmarkPermutation3Order1(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b3.Permutation3ParamOrder1)
}

func BenchmarkPermutation3Order2(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b3.Permutation3ParamOrder2)
}

func BenchmarkPermutation3Order3(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b3.Permutation3ParamOrder3)
}

func BenchmarkPermutation3Order4(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b3.Permutation3ParamOrder4)
}

func BenchmarkPermutation3Order5(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b3.Permutation3ParamOrder5)
}
