package leadingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/table"
)

// BenchmarkLeadingZeros32Calibrate-8   	1000000000	         2.43 ns/op
// BenchmarkLeadingZeros32Naive-8       	300000000	         5.10 ns/op
// BenchmarkLeadingZeros32Table-8       	300000000	         4.57 ns/op <- Best
// BenchmarkLeadingZeros32Stdlib-8      	300000000	         5.40 ns/op

var uint32Supplier = randomsupplier.NewUint32()

func BenchmarkLeadingZeros32CalibrateSupplier(b *testing.B) {
	benchmark.Uint32Supplier(b, uint32Supplier.Next)
}

func BenchmarkLeadingZeros32CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint32ToIntFunc(b, uint32Supplier.Next)
}

func BenchmarkLeadingZeros32Naive(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, naive.LeadingZeros32)
}

func BenchmarkLeadingZeros32Table(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, table.LeadingZeros32)
}

func BenchmarkLeadingZeros32Stdlib(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, stdlib.LeadingZeros32)
}
