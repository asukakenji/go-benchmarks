package b2_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/permutation/b2"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkUintSliceGenerator(b *testing.B) {
	bcommon.Reset()
	benchmark.CalibrateUintSliceGenerator(b, bcommon.NextSlice)
}

// BenchmarkPermutation1Order0-8   	    3000	    440636 ns/op
// BenchmarkPermutation1Order1-8   	    3000	    444581 ns/op
// BenchmarkPermutation1Order2-8   	    3000	    437564 ns/op <- Best
// BenchmarkPermutation1Order3-8   	    3000	    433765 ns/op <- Best
// BenchmarkPermutation1Order4-8   	    3000	    442313 ns/op
// BenchmarkPermutation1Order5-8   	    3000	    432704 ns/op <- Best

func BenchmarkPermutation1Order0(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation1ParamOrder0)
}

func BenchmarkPermutation1Order1(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation1ParamOrder1)
}

func BenchmarkPermutation1Order2(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation1ParamOrder2)
}

func BenchmarkPermutation1Order3(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation1ParamOrder3)
}

func BenchmarkPermutation1Order4(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation1ParamOrder4)
}

func BenchmarkPermutation1Order5(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation1ParamOrder5)
}
