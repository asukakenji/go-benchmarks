package trailingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/trailingzeros/stdlib"
)

var uint16Supplier = randomsupplier.NewUint16()

func BenchmarkLeadingZeros16CalibrateSupplier(b *testing.B) {
	benchmark.Uint16Supplier(b, uint16Supplier.Next)
}

func BenchmarkLeadingZeros16CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint16ToIntFunc(b, uint16Supplier.Next)
}

func BenchmarkTrailingZeros16Naive(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, naive.TrailingZeros16)
}

func BenchmarkTrailingZeros16Stdlib(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, stdlib.TrailingZeros16)
}
