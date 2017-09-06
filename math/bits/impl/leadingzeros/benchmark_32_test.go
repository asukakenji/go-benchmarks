package leadingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz1"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz10"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz10b"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz1a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz3"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz4"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz5"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/table"
)

// BenchmarkLeadingZeros32CalibrateSupplier-8      	500000000	         3.52 ns/op
// BenchmarkLeadingZeros32CalibrateBenchmarker-8   	300000000	         4.56 ns/op
// BenchmarkLeadingZeros32Naive-8                  	200000000	         8.38 ns/op
// BenchmarkLeadingZeros32Table-8                  	300000000	         5.37 ns/op <- Best
// BenchmarkLeadingZeros32Stdlib-8                 	200000000	         6.38 ns/op <- Best
// BenchmarkLeadingZeros32Nlz1-8                   	100000000	        10.1 ns/op
// BenchmarkLeadingZeros32Nlz1a-8                  	200000000	         9.20 ns/op
// BenchmarkLeadingZeros32Nlz2-8                   	200000000	         7.15 ns/op <- Best (Non-Table)
// BenchmarkLeadingZeros32Nlz2a-8                  	100000000	        13.5 ns/op
// BenchmarkLeadingZeros32Nlz3-8                   	100000000	        22.3 ns/op
// BenchmarkLeadingZeros32Nlz4-8                   	200000000	         9.38 ns/op
// BenchmarkLeadingZeros32Nlz5-8                   	200000000	         8.18 ns/op
// BenchmarkLeadingZeros32Nlz10-8                  	200000000	         6.72 ns/op <- Best
// BenchmarkLeadingZeros32Nlz10NoMultiply-8        	200000000	         7.88 ns/op
// BenchmarkLeadingZeros32Nlz10b-8                 	200000000	         6.92 ns/op <- Best
// BenchmarkLeadingZeros32Nlz10bNoMultiply-8       	200000000	         7.74 ns/op

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

func BenchmarkLeadingZeros32Nlz1(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, nlz1.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz1a(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, nlz1a.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz2(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, nlz2.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz2a(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, nlz2a.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz3(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, nlz3.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz4(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, nlz4.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz5(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, nlz5.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz10(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, nlz10.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz10NoMultiply(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, nlz10.LeadingZeros32NoMultiply)
}

func BenchmarkLeadingZeros32Nlz10b(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, nlz10b.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz10bNoMultiply(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, nlz10b.LeadingZeros32NoMultiply)
}
