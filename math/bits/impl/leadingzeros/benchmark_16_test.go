package leadingzeros_test

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

// BenchmarkLeadingZeros16CalibrateSupplier-8      	500000000	         3.46 ns/op
// BenchmarkLeadingZeros16CalibrateBenchmarker-8   	300000000	         4.87 ns/op
// BenchmarkLeadingZeros16Naive-8                  	100000000	        13.0 ns/op
// BenchmarkLeadingZeros16Table-8                  	300000000	         5.57 ns/op <- Best
// BenchmarkLeadingZeros16Stdlib-8                 	200000000	         6.43 ns/op <- Best
// BenchmarkLeadingZeros16Nlz1-8                   	100000000	        11.8 ns/op
// BenchmarkLeadingZeros16Nlz1a-8                  	100000000	        11.0 ns/op
// BenchmarkLeadingZeros16Nlz2-8                   	200000000	         8.62 ns/op
// BenchmarkLeadingZeros16Nlz2a-8                  	100000000	        14.9 ns/op
// BenchmarkLeadingZeros16Nlz3-8                   	100000000	        17.8 ns/op
// BenchmarkLeadingZeros16Nlz5-8                   	200000000	         7.48 ns/op <- Best (Non-Table)

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

func BenchmarkLeadingZeros16Nlz1(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, nlz1.LeadingZeros16)
}

func BenchmarkLeadingZeros16Nlz1a(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, nlz1a.LeadingZeros16)
}

func BenchmarkLeadingZeros16Nlz2(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, nlz2.LeadingZeros16)
}

func BenchmarkLeadingZeros16Nlz2a(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, nlz2a.LeadingZeros16)
}

func BenchmarkLeadingZeros16Nlz3(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, nlz3.LeadingZeros16)
}

func BenchmarkLeadingZeros16Nlz5(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, nlz5.LeadingZeros16)
}
