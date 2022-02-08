package b1_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b1"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkPermutation1Order0_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder0[uint], 8)
}

func BenchmarkPermutation1Order1_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder1[uint], 8)
}

func BenchmarkPermutation1Order2_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder2[uint], 8)
}

func BenchmarkPermutation1Order3_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder3[uint], 8)
}

func BenchmarkPermutation1Order4_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder4[uint], 8)
}

func BenchmarkPermutation1Order5_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder5[uint], 8)
}

// ---

func BenchmarkPermutation1Order0_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder0[uint], 9)
}

func BenchmarkPermutation1Order1_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder1[uint], 9)
}

func BenchmarkPermutation1Order2_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder2[uint], 9)
}

func BenchmarkPermutation1Order3_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder3[uint], 9)
}

func BenchmarkPermutation1Order4_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder4[uint], 9)
}

func BenchmarkPermutation1Order5_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder5[uint], 9)
}

// ---

func BenchmarkPermutation1Order0_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder0[uint], 10)
}

func BenchmarkPermutation1Order1_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder1[uint], 10)
}

func BenchmarkPermutation1Order2_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder2[uint], 10)
}

func BenchmarkPermutation1Order3_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder3[uint], 10)
}

func BenchmarkPermutation1Order4_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder4[uint], 10)
}

func BenchmarkPermutation1Order5_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b1.Permutation1ParamOrder5[uint], 10)
}
