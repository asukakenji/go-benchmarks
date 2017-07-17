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

// BenchmarkPermutation1Order0-8   	    3000	    440315 ns/op
// BenchmarkPermutation1Order1-8   	    3000	    444352 ns/op
// BenchmarkPermutation1Order2-8   	    3000	    439026 ns/op <- Best
// BenchmarkPermutation1Order3-8   	    3000	    434343 ns/op <- Best
// BenchmarkPermutation1Order4-8   	    3000	    444893 ns/op
// BenchmarkPermutation1Order5-8   	    3000	    434416 ns/op <- Best

func BenchmarkPermutation1Order0(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1.Permutation1ParamOrder0)
}

func BenchmarkPermutation1Order1(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1.Permutation1ParamOrder1)
}

func BenchmarkPermutation1Order2(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1.Permutation1ParamOrder2)
}

func BenchmarkPermutation1Order3(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1.Permutation1ParamOrder3)
}

func BenchmarkPermutation1Order4(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1.Permutation1ParamOrder4)
}

func BenchmarkPermutation1Order5(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1.Permutation1ParamOrder5)
}
