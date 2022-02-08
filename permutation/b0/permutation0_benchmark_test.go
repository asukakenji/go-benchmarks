package b0_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b0"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkPermutation0Order0_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder0[uint], 8)
}

func BenchmarkPermutation0Order1_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder1[uint], 8)
}

func BenchmarkPermutation0Order2_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder2[uint], 8)
}

func BenchmarkPermutation0Order3_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder3[uint], 8)
}

func BenchmarkPermutation0Order4_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder4[uint], 8)
}

func BenchmarkPermutation0Order5_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder5[uint], 8)
}

// ---

func BenchmarkPermutation0Order0_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder0[uint], 9)
}

func BenchmarkPermutation0Order1_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder1[uint], 9)
}

func BenchmarkPermutation0Order2_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder2[uint], 9)
}

func BenchmarkPermutation0Order3_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder3[uint], 9)
}

func BenchmarkPermutation0Order4_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder4[uint], 9)
}

func BenchmarkPermutation0Order5_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder5[uint], 9)
}

// ---

func BenchmarkPermutation0Order0_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder0[uint], 10)
}

func BenchmarkPermutation0Order1_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder1[uint], 10)
}

func BenchmarkPermutation0Order2_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder2[uint], 10)
}

func BenchmarkPermutation0Order3_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder3[uint], 10)
}

func BenchmarkPermutation0Order4_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder4[uint], 10)
}

func BenchmarkPermutation0Order5_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b0.Permutation0ParamOrder5[uint], 10)
}
