package leadingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/table"
)

// BenchmarkLeadingZeros64Calibrate-8   	1000000000	         2.38 ns/op
// BenchmarkLeadingZeros64Naive-8       	300000000	         4.69 ns/op
// BenchmarkLeadingZeros64Table-8       	300000000	         4.47 ns/op <- Best
// BenchmarkLeadingZeros64Stdlib-8      	300000000	         5.47 ns/op

var uint64Supplier = randomsupplier.NewUint64()

func BenchmarkLeadingZeros64CalibrateSupplier(b *testing.B) {
	benchmark.Uint64Supplier(b, uint64Supplier.Next)
}

func BenchmarkLeadingZeros64CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint64ToIntFunc(b, uint64Supplier.Next)
}

func BenchmarkLeadingZeros64Naive(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, naive.LeadingZeros64)
}

func BenchmarkLeadingZeros64Table(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, table.LeadingZeros64)
}

func BenchmarkLeadingZeros64Stdlib(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, stdlib.LeadingZeros64)
}
