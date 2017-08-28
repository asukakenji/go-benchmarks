package trailingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/stdlib"
)

func BenchmarkTrailingZeros32Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUint32WithRandom(b)
}

func BenchmarkTrailingZeros32Naive(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, naive.TrailingZeros32)
}

func BenchmarkTrailingZeros32Stdlib(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, stdlib.TrailingZeros32)
}
