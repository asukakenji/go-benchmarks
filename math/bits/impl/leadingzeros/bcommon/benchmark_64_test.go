package bcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz1"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz1a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz3"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz5"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/table"
)

// BenchmarkLeadingZeros64CalibrateSupplier-8      	500000000	         3.58 ns/op
// BenchmarkLeadingZeros64CalibrateBenchmarker-8   	300000000	         4.71 ns/op
// BenchmarkLeadingZeros64Naive-8                  	200000000	         6.41 ns/op <- Best (Non-Table)
// BenchmarkLeadingZeros64Table-8                  	300000000	         5.61 ns/op <- Best
// BenchmarkLeadingZeros64Stdlib-8                 	200000000	         7.13 ns/op
// BenchmarkLeadingZeros64Nlz1-8                   	100000000	        10.3 ns/op
// BenchmarkLeadingZeros64Nlz1a-8                  	200000000	         6.87 ns/op
// BenchmarkLeadingZeros64Nlz2-8                   	200000000	         7.47 ns/op
// BenchmarkLeadingZeros64Nlz2a-8                  	100000000	        13.6 ns/op
// BenchmarkLeadingZeros64Nlz3-8                   	50000000	        33.1 ns/op
// BenchmarkLeadingZeros64Nlz5-8                   	200000000	         9.19 ns/op

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

func BenchmarkLeadingZeros64Nlz1(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, nlz1.LeadingZeros64)
}

func BenchmarkLeadingZeros64Nlz1a(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, nlz1a.LeadingZeros64)
}

func BenchmarkLeadingZeros64Nlz2(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, nlz2.LeadingZeros64)
}

func BenchmarkLeadingZeros64Nlz2a(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, nlz2a.LeadingZeros64)
}

func BenchmarkLeadingZeros64Nlz3(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, nlz3.LeadingZeros64)
}

func BenchmarkLeadingZeros64Nlz5(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, nlz5.LeadingZeros64)
}
