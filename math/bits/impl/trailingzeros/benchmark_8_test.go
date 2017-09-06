package trailingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/stdlib"
)

var uint8Supplier = randomsupplier.NewUint8()

func BenchmarkLeadingZeros8CalibrateSupplier(b *testing.B) {
	benchmark.Uint8Supplier(b, uint8Supplier.Next)
}

func BenchmarkLeadingZeros8CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint8ToIntFunc(b, uint8Supplier.Next)
}

func BenchmarkTrailingZeros8Naive(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, naive.TrailingZeros8)
}

func BenchmarkTrailingZeros8Stdlib(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, stdlib.TrailingZeros8)
}
