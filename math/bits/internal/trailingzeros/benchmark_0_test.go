package trailingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/stdlib"
)

func BenchmarkTrailingZeros0Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUintWithRandom(b)
}

func BenchmarkTrailingZeros0Naive(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, naive.TrailingZeros)
}

func BenchmarkTrailingZeros0Stdlib(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, stdlib.TrailingZeros)
}
