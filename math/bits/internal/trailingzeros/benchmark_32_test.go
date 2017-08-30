package trailingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/stdlib"
)

var uint32Supplier = randomsupplier.NewUint32()

func BenchmarkLeadingZeros32CalibrateSupplier(b *testing.B) {
	benchmark.Uint32Supplier(b, uint32Supplier.Next)
}

func BenchmarkLeadingZeros32CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint32ToIntFunc(b, uint32Supplier.Next)
}

func BenchmarkTrailingZeros32Naive(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, naive.TrailingZeros32)
}

func BenchmarkTrailingZeros32Stdlib(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, stdlib.TrailingZeros32)
}
