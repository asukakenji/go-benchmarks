package trailingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/trailingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/trailingzeros/stdlib"
)

var uintSupplier = randomsupplier.NewUint()

func BenchmarkLeadingZeros0CalibrateSupplier(b *testing.B) {
	benchmark.UintSupplier(b, uintSupplier.Next)
}

func BenchmarkLeadingZeros0CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUintToIntFunc(b, uintSupplier.Next)
}

func BenchmarkTrailingZeros0Naive(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, naive.TrailingZeros)
}

func BenchmarkTrailingZeros0Stdlib(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, stdlib.TrailingZeros)
}
