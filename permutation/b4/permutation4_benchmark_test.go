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

// BenchmarkPermutation4Order0-8   	    3000	    440636 ns/op
// BenchmarkPermutation4Order1-8   	    3000	    444581 ns/op
// BenchmarkPermutation4Order2-8   	    3000	    437564 ns/op <- Best
// BenchmarkPermutation4Order3-8   	    3000	    433765 ns/op <- Best
// BenchmarkPermutation4Order4-8   	    3000	    442313 ns/op
// BenchmarkPermutation4Order5-8   	    3000	    432704 ns/op <- Best

func BenchmarkPermutation4Order0(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b4.Permutation4ParamOrder0)
}

func BenchmarkPermutation4Order1(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b4.Permutation4ParamOrder1)
}

func BenchmarkPermutation4Order2(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b4.Permutation4ParamOrder2)
}

func BenchmarkPermutation4Order3(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b4.Permutation4ParamOrder3)
}

func BenchmarkPermutation4Order4(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b4.Permutation4ParamOrder4)
}

func BenchmarkPermutation4Order5(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b4.Permutation4ParamOrder5)
}
