package trailingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/stdlib"
)

func BenchmarkTrailingZeros8Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUint8WithRandom(b)
}

func BenchmarkTrailingZeros8Naive(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, naive.TrailingZeros8)
}

func BenchmarkTrailingZeros8Stdlib(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, stdlib.TrailingZeros8)
}
