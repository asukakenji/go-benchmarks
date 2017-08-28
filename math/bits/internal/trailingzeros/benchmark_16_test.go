package trailingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/stdlib"
)

func BenchmarkTrailingZeros16Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUint16WithRandom(b)
}

func BenchmarkTrailingZeros16Naive(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, naive.TrailingZeros16)
}

func BenchmarkTrailingZeros16Stdlib(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, stdlib.TrailingZeros16)
}
