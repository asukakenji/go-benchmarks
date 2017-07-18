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

// BenchmarkPermutation2Order00-8   	    3000	    460672 ns/op
// BenchmarkPermutation2Order01-8   	    3000	    463003 ns/op
// BenchmarkPermutation2Order02-8   	    3000	    461831 ns/op
// BenchmarkPermutation2Order03-8   	    3000	    461485 ns/op
// BenchmarkPermutation2Order10-8   	    3000	    466685 ns/op
// BenchmarkPermutation2Order11-8   	    3000	    466419 ns/op
// BenchmarkPermutation2Order12-8   	    3000	    468907 ns/op
// BenchmarkPermutation2Order13-8   	    3000	    466987 ns/op
// BenchmarkPermutation2Order20-8   	    3000	    454632 ns/op
// BenchmarkPermutation2Order21-8   	    3000	    456961 ns/op
// BenchmarkPermutation2Order22-8   	    3000	    455437 ns/op
// BenchmarkPermutation2Order23-8   	    3000	    452663 ns/op <- Best
// BenchmarkPermutation2Order30-8   	    3000	    457167 ns/op
// BenchmarkPermutation2Order31-8   	    3000	    457754 ns/op
// BenchmarkPermutation2Order32-8   	    3000	    459530 ns/op
// BenchmarkPermutation2Order33-8   	    3000	    454949 ns/op <- Best
// BenchmarkPermutation2Order40-8   	    3000	    467237 ns/op
// BenchmarkPermutation2Order41-8   	    3000	    465731 ns/op
// BenchmarkPermutation2Order42-8   	    3000	    466926 ns/op
// BenchmarkPermutation2Order43-8   	    3000	    462871 ns/op
// BenchmarkPermutation2Order50-8   	    3000	    454398 ns/op <- Best
// BenchmarkPermutation2Order51-8   	    3000	    454527 ns/op
// BenchmarkPermutation2Order52-8   	    3000	    457110 ns/op
// BenchmarkPermutation2Order53-8   	    3000	    456651 ns/op

func BenchmarkPermutation2Order00(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder00)
}

func BenchmarkPermutation2Order01(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder01)
}

func BenchmarkPermutation2Order02(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder02)
}

func BenchmarkPermutation2Order03(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder03)
}

func BenchmarkPermutation2Order10(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder10)
}

func BenchmarkPermutation2Order11(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder11)
}

func BenchmarkPermutation2Order12(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder12)
}

func BenchmarkPermutation2Order13(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder13)
}

func BenchmarkPermutation2Order20(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder20)
}

func BenchmarkPermutation2Order21(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder21)
}

func BenchmarkPermutation2Order22(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder22)
}

func BenchmarkPermutation2Order23(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder23)
}

func BenchmarkPermutation2Order30(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder30)
}

func BenchmarkPermutation2Order31(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder31)
}

func BenchmarkPermutation2Order32(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder32)
}

func BenchmarkPermutation2Order33(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder33)
}

func BenchmarkPermutation2Order40(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder40)
}

func BenchmarkPermutation2Order41(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder41)
}

func BenchmarkPermutation2Order42(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder42)
}

func BenchmarkPermutation2Order43(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder43)
}

func BenchmarkPermutation2Order50(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder50)
}

func BenchmarkPermutation2Order51(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder51)
}

func BenchmarkPermutation2Order52(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder52)
}

func BenchmarkPermutation2Order53(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b2.Permutation2ParamOrder53)
}
