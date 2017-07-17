package b1b_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/permutation/b1b"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkUintSliceGenerator(b *testing.B) {
	bcommon.Reset()
	benchmark.CalibrateUintSliceGenerator(b, bcommon.NextSlice)
}

// BenchmarkPermutation1Order00-8   	    3000	    460672 ns/op
// BenchmarkPermutation1Order01-8   	    3000	    463003 ns/op
// BenchmarkPermutation1Order02-8   	    3000	    461831 ns/op
// BenchmarkPermutation1Order03-8   	    3000	    461485 ns/op
// BenchmarkPermutation1Order10-8   	    3000	    466685 ns/op
// BenchmarkPermutation1Order11-8   	    3000	    466419 ns/op
// BenchmarkPermutation1Order12-8   	    3000	    468907 ns/op
// BenchmarkPermutation1Order13-8   	    3000	    466987 ns/op
// BenchmarkPermutation1Order20-8   	    3000	    454632 ns/op
// BenchmarkPermutation1Order21-8   	    3000	    456961 ns/op
// BenchmarkPermutation1Order22-8   	    3000	    455437 ns/op
// BenchmarkPermutation1Order23-8   	    3000	    452663 ns/op <- Best
// BenchmarkPermutation1Order30-8   	    3000	    457167 ns/op
// BenchmarkPermutation1Order31-8   	    3000	    457754 ns/op
// BenchmarkPermutation1Order32-8   	    3000	    459530 ns/op
// BenchmarkPermutation1Order33-8   	    3000	    454949 ns/op <- Best
// BenchmarkPermutation1Order40-8   	    3000	    467237 ns/op
// BenchmarkPermutation1Order41-8   	    3000	    465731 ns/op
// BenchmarkPermutation1Order42-8   	    3000	    466926 ns/op
// BenchmarkPermutation1Order43-8   	    3000	    462871 ns/op
// BenchmarkPermutation1Order50-8   	    3000	    454398 ns/op <- Best
// BenchmarkPermutation1Order51-8   	    3000	    454527 ns/op
// BenchmarkPermutation1Order52-8   	    3000	    457110 ns/op
// BenchmarkPermutation1Order53-8   	    3000	    456651 ns/op

func BenchmarkPermutation1Order00(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder00)
}

func BenchmarkPermutation1Order01(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder01)
}

func BenchmarkPermutation1Order02(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder02)
}

func BenchmarkPermutation1Order03(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder03)
}

func BenchmarkPermutation1Order10(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder10)
}

func BenchmarkPermutation1Order11(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder11)
}

func BenchmarkPermutation1Order12(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder12)
}

func BenchmarkPermutation1Order13(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder13)
}

func BenchmarkPermutation1Order20(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder20)
}

func BenchmarkPermutation1Order21(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder21)
}

func BenchmarkPermutation1Order22(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder22)
}

func BenchmarkPermutation1Order23(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder23)
}

func BenchmarkPermutation1Order30(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder30)
}

func BenchmarkPermutation1Order31(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder31)
}

func BenchmarkPermutation1Order32(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder32)
}

func BenchmarkPermutation1Order33(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder33)
}

func BenchmarkPermutation1Order40(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder40)
}

func BenchmarkPermutation1Order41(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder41)
}

func BenchmarkPermutation1Order42(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder42)
}

func BenchmarkPermutation1Order43(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder43)
}

func BenchmarkPermutation1Order50(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder50)
}

func BenchmarkPermutation1Order51(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder51)
}

func BenchmarkPermutation1Order52(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder52)
}

func BenchmarkPermutation1Order53(b *testing.B) {
	bcommon.Reset()
	benchmark.UintSliceGenerator(b, bcommon.NextSlice, b1b.Permutation1ParamOrder53)
}
