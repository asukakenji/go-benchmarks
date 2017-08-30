package leadingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/table"
)

// BenchmarkLeadingZeros8Calibrate-8   	1000000000	         2.21 ns/op
// BenchmarkLeadingZeros8Naive-8       	100000000	        10.0 ns/op
// BenchmarkLeadingZeros8Table-8       	300000000	         4.27 ns/op <- Best
// BenchmarkLeadingZeros8Stdlib-8      	300000000	         4.49 ns/op

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
