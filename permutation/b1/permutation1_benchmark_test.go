package b1_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/permutation/b1"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkUintSliceGenerator(b *testing.B) {
	bcommon.Reset()
	benchmark.CalibrateUintSliceGenerator(b, bcommon.NextSlice)
}

// BenchmarkPermutation1Order0-8   	    3000	    445322 ns/op
// BenchmarkPermutation1Order1-8   	    3000	    448333 ns/op
// BenchmarkPermutation1Order2-8   	    3000	    444541 ns/op
// BenchmarkPermutation1Order3-8   	    3000	    436832 ns/op <- Best
// BenchmarkPermutation1Order4-8   	    3000	    443842 ns/op
// BenchmarkPermutation1Order5-8   	    3000	    438226 ns/op <- Best

func BenchmarkPermutation1Order0(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1.Permutation1ParamOrder0[uint])
}

func BenchmarkPermutation1Order1(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1.Permutation1ParamOrder1[uint])
}

func BenchmarkPermutation1Order2(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1.Permutation1ParamOrder2[uint])
}

func BenchmarkPermutation1Order3(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1.Permutation1ParamOrder3[uint])
}

func BenchmarkPermutation1Order4(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1.Permutation1ParamOrder4[uint])
}

func BenchmarkPermutation1Order5(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1.Permutation1ParamOrder5[uint])
}
