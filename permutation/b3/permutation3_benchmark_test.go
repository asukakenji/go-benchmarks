package b3_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b3"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkPermutation3Order0_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder0[uint], 8)
}

func BenchmarkPermutation3Order1_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder1[uint], 8)
}

func BenchmarkPermutation3Order2_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder2[uint], 8)
}

func BenchmarkPermutation3Order3_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder3[uint], 8)
}

func BenchmarkPermutation3Order4_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder4[uint], 8)
}

func BenchmarkPermutation3Order5_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder5[uint], 8)
}

// ---

func BenchmarkPermutation3Order0_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder0[uint], 9)
}

func BenchmarkPermutation3Order1_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder1[uint], 9)
}

func BenchmarkPermutation3Order2_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder2[uint], 9)
}

func BenchmarkPermutation3Order3_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder3[uint], 9)
}

func BenchmarkPermutation3Order4_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder4[uint], 9)
}

func BenchmarkPermutation3Order5_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder5[uint], 9)
}

// ---

func BenchmarkPermutation3Order0_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder0[uint], 10)
}

func BenchmarkPermutation3Order1_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder1[uint], 10)
}

func BenchmarkPermutation3Order2_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder2[uint], 10)
}

func BenchmarkPermutation3Order3_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder3[uint], 10)
}

func BenchmarkPermutation3Order4_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder4[uint], 10)
}

func BenchmarkPermutation3Order5_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b3.Permutation3ParamOrder5[uint], 10)
}
