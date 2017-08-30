package trailingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/stdlib"
)

var uint64Supplier = randomsupplier.NewUint64()

func BenchmarkLeadingZeros64CalibrateSupplier(b *testing.B) {
	benchmark.Uint64Supplier(b, uint64Supplier.Next)
}

func BenchmarkLeadingZeros64CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint64ToIntFunc(b, uint64Supplier.Next)
}

func BenchmarkTrailingZeros64Naive(b *testing.B) {
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, naive.TrailingZeros64)
}

func BenchmarkTrailingZeros64Stdlib(b *testing.B) {
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, stdlib.TrailingZeros64)
}
