package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount16CalibrateSupplier-8               	500000000	         3.42 ns/op
// BenchmarkOnesCount16CalibrateBenchmarker-8            	300000000	         4.52 ns/op
// BenchmarkOnesCount16Naive-8                           	30000000	        57.2 ns/op
// BenchmarkOnesCount16Table-8                           	300000000	         5.58 ns/op <- Best
// BenchmarkOnesCount16Stdlib-8                          	300000000	         5.75 ns/op <- Best
// BenchmarkOnesCount16Pop0-8                            	200000000	         7.07 ns/op
// BenchmarkOnesCount16Pop1-8                            	200000000	         6.63 ns/op
// BenchmarkOnesCount16Pop1A-8                           	200000000	         6.53 ns/op
// BenchmarkOnesCount16Reset-8                           	100000000	        17.8 ns/op
// BenchmarkOnesCount16Subtract-8                        	100000000	        17.4 ns/op

var uint16Supplier = randomsupplier.NewUint16()

func BenchmarkOnesCount16CalibrateSupplier(b *testing.B) {
	benchmark.Uint16Supplier(b, uint16Supplier.Next)
}

func BenchmarkOnesCount16CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint16ToIntFunc(b, uint16Supplier.Next)
}

func BenchmarkOnesCount16Naive(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, naive.OnesCount16)
}

func BenchmarkOnesCount16Table(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, table.OnesCount16)
}

func BenchmarkOnesCount16Stdlib(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, stdlib.OnesCount16)
}

func BenchmarkOnesCount16Pop0(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, pop0.OnesCount16)
}

func BenchmarkOnesCount16Pop1(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, pop1.OnesCount16)
}

func BenchmarkOnesCount16Pop1A(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, pop1a.OnesCount16)
}

func BenchmarkOnesCount16Reset(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, reset.OnesCount16)
}

func BenchmarkOnesCount16Subtract(b *testing.B) {
	uint16Supplier.Reset()
	benchmark.Uint16ToIntFunc(b, uint16Supplier.Next, subtract.OnesCount16)
}
