package trailingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/stdlib"
)

func BenchmarkTrailingZeros64Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUint64WithRandom(b)
}

func BenchmarkTrailingZeros64Naive(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, naive.TrailingZeros64)
}

func BenchmarkTrailingZeros64Stdlib(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, stdlib.TrailingZeros64)
}
