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

// BenchmarkLeadingZeros8CalibrateSupplier-8      	500000000	         3.68 ns/op
// BenchmarkLeadingZeros8CalibrateBenchmarker-8   	300000000	         4.96 ns/op
// BenchmarkLeadingZeros8Naive-8                  	100000000	        12.3 ns/op
// BenchmarkLeadingZeros8Table-8                  	300000000	         5.41 ns/op <- Best
// BenchmarkLeadingZeros8Stdlib-8                 	300000000	         5.54 ns/op <- Best
// BenchmarkLeadingZeros8Nlz1-8                   	100000000	        12.4 ns/op
// BenchmarkLeadingZeros8Nlz1a-8                  	100000000	        12.7 ns/op
// BenchmarkLeadingZeros8Nlz2-8                   	100000000	        11.3 ns/op
// BenchmarkLeadingZeros8Nlz2a-8                  	100000000	        14.0 ns/op
// BenchmarkLeadingZeros8Nlz3-8                   	100000000	        14.8 ns/op
// BenchmarkLeadingZeros8Nlz5-8                   	200000000	         6.68 ns/op <- Best (Non-Table)

var uint8Supplier = randomsupplier.NewUint8()

func BenchmarkLeadingZeros8CalibrateSupplier(b *testing.B) {
	benchmark.Uint8Supplier(b, uint8Supplier.Next)
}

func BenchmarkLeadingZeros8CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint8ToIntFunc(b, uint8Supplier.Next)
}

func BenchmarkLeadingZeros8Naive(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, naive.LeadingZeros8)
}

func BenchmarkLeadingZeros8Table(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, table.LeadingZeros8)
}

func BenchmarkLeadingZeros8Stdlib(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, stdlib.LeadingZeros8)
}

func BenchmarkLeadingZeros8Nlz1(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, nlz1.LeadingZeros8)
}

func BenchmarkLeadingZeros8Nlz1a(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, nlz1a.LeadingZeros8)
}

func BenchmarkLeadingZeros8Nlz2(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, nlz2.LeadingZeros8)
}

func BenchmarkLeadingZeros8Nlz2a(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, nlz2a.LeadingZeros8)
}

func BenchmarkLeadingZeros8Nlz3(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, nlz3.LeadingZeros8)
}

func BenchmarkLeadingZeros8Nlz5(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, nlz5.LeadingZeros8)
}
