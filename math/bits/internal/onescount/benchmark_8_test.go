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

// BenchmarkOnesCount8CalibrateSupplier-8               	500000000	         3.56 ns/op
// BenchmarkOnesCount8CalibrateBenchmarker-8            	300000000	         5.02 ns/op
// BenchmarkOnesCount8Naive-8                           	50000000	        31.7 ns/op
// BenchmarkOnesCount8Table-8                           	300000000	         5.42 ns/op <- Best
// BenchmarkOnesCount8Stdlib-8                          	300000000	         5.41 ns/op <- Best
// BenchmarkOnesCount8Pop0-8                            	200000000	         6.84 ns/op
// BenchmarkOnesCount8Pop1-8                            	200000000	         6.32 ns/op
// BenchmarkOnesCount8Pop1A-8                           	200000000	         6.33 ns/op
// BenchmarkOnesCount8Reset-8                           	100000000	        15.4 ns/op
// BenchmarkOnesCount8Subtract-8                        	100000000	        13.4 ns/op

var uint8Supplier = randomsupplier.NewUint8()

func BenchmarkOnesCount8CalibrateSupplier(b *testing.B) {
	benchmark.Uint8Supplier(b, uint8Supplier.Next)
}

func BenchmarkOnesCount8CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint8ToIntFunc(b, uint8Supplier.Next)
}

func BenchmarkOnesCount8Naive(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, naive.OnesCount8)
}

func BenchmarkOnesCount8Table(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, table.OnesCount8)
}

func BenchmarkOnesCount8Stdlib(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, stdlib.OnesCount8)
}

func BenchmarkOnesCount8Pop0(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, pop0.OnesCount8)
}

func BenchmarkOnesCount8Pop1(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, pop1.OnesCount8)
}

func BenchmarkOnesCount8Pop1A(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, pop1a.OnesCount8)
}

func BenchmarkOnesCount8Reset(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, reset.OnesCount8)
}

func BenchmarkOnesCount8Subtract(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, subtract.OnesCount8)
}
