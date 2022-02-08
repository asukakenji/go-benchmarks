package b5_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b5"
	"github.com/asukakenji/go-benchmarks/permutation/bcommon"
)

func BenchmarkPermutation5Inc_8(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b5.Permutation5Inc[uint], 8)
}

func BenchmarkPermutation5Inc_9(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b5.Permutation5Inc[uint], 9)
}

func BenchmarkPermutation5Inc_10(b *testing.B) {
	bcommon.BenchmarkPermutation(b, b5.Permutation5Inc[uint], 10)
}
