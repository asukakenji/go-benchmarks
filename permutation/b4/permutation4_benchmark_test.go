package b4_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b4"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkPermutation4Order0_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder0[uint], 8)
}

func BenchmarkPermutation4Order1_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder1[uint], 8)
}

func BenchmarkPermutation4Order2_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder2[uint], 8)
}

func BenchmarkPermutation4Order3_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder3[uint], 8)
}

func BenchmarkPermutation4Order4_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder4[uint], 8)
}

func BenchmarkPermutation4Order5_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder5[uint], 8)
}

// ---

func BenchmarkPermutation4Order0_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder0[uint], 9)
}

func BenchmarkPermutation4Order1_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder1[uint], 9)
}

func BenchmarkPermutation4Order2_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder2[uint], 9)
}

func BenchmarkPermutation4Order3_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder3[uint], 9)
}

func BenchmarkPermutation4Order4_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder4[uint], 9)
}

func BenchmarkPermutation4Order5_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder5[uint], 9)
}

// ---

func BenchmarkPermutation4Order0_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder0[uint], 10)
}

func BenchmarkPermutation4Order1_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder1[uint], 10)
}

func BenchmarkPermutation4Order2_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder2[uint], 10)
}

func BenchmarkPermutation4Order3_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder3[uint], 10)
}

func BenchmarkPermutation4Order4_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder4[uint], 10)
}

func BenchmarkPermutation4Order5_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b4.Permutation4ParamOrder5[uint], 10)
}
