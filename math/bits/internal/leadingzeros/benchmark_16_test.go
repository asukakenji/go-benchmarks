package leadingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/table"
)

// BenchmarkLeadingZeros16Calibrate-8   	1000000000	         2.36 ns/op
// BenchmarkLeadingZeros16Naive-8       	200000000	         9.63 ns/op
// BenchmarkLeadingZeros16Table-8       	300000000	         4.49 ns/op <- Best
// BenchmarkLeadingZeros16Stdlib-8      	300000000	         5.14 ns/op

var uint16Supplier = randomsupplier.NewUint16()

func BenchmarkLeadingZeros16CalibrateSupplier(b *testing.B) {
	benchmark.Uint16Supplier(b, uint16Supplier.Next)
}

func BenchmarkLeadingZeros16CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint16ToIntFunc(b, uint16Supplier.Next)
}

func BenchmarkLeadingZeros16Naive(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, naive.LeadingZeros16)
}

func BenchmarkLeadingZeros16Table(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, table.LeadingZeros16)
}

func BenchmarkLeadingZeros16Stdlib(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, stdlib.LeadingZeros16)
}
