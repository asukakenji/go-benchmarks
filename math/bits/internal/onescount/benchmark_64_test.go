package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/gccbuiltin"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount64CalibrateSupplier-8               	500000000	         3.58 ns/op
// BenchmarkOnesCount64CalibrateBenchmarker-8            	300000000	         4.84 ns/op
// BenchmarkOnesCount64Naive-8                           	10000000	       186 ns/op
// BenchmarkOnesCount64Table-8                           	200000000	         9.51 ns/op
// BenchmarkOnesCount64Stdlib-8                          	200000000	         7.85 ns/op
// BenchmarkOnesCount64Pop0-8                            	200000000	         8.72 ns/op
// BenchmarkOnesCount64Pop1-8                            	200000000	         7.81 ns/op
// BenchmarkOnesCount64Pop1Alt-8                         	200000000	         6.89 ns/op <- Best
// BenchmarkOnesCount64Reset-8                           	50000000	        34.9 ns/op
// BenchmarkOnesCount64Subtract-8                        	30000000	        45.9 ns/op
// BenchmarkOnesCount64CallGCC-8                         	20000000	        66.5 ns/op
// BenchmarkOnesCount64Pop3-8                            	200000000	         7.19 ns/op <- Best
// BenchmarkOnesCount64Pop5-8                            	 5000000	       315 ns/op
// BenchmarkOnesCount64Hakmem-8                          	200000000	         8.08 ns/op
// BenchmarkOnesCount64Asm-8                             	200000000	         7.07 ns/op <- Best

var uint64Supplier = randomsupplier.NewUint64()

func BenchmarkOnesCount64CalibrateSupplier(b *testing.B) {
	benchmark.Uint64Supplier(b, uint64Supplier.Next)
}

func BenchmarkOnesCount64CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint64ToIntFunc(b, uint64Supplier.Next)
}

func BenchmarkOnesCount64Naive(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, naive.OnesCount64)
}

func BenchmarkOnesCount64Table(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, table.OnesCount64)
}

func BenchmarkOnesCount64Stdlib(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, stdlib.OnesCount64)
}

func BenchmarkOnesCount64Pop0(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, pop0.OnesCount64)
}

func BenchmarkOnesCount64Pop1(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, pop1.OnesCount64)
}

func BenchmarkOnesCount64Pop1Alt(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, pop1a.OnesCount64)
}

func BenchmarkOnesCount64Reset(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, reset.OnesCount64)
}

func BenchmarkOnesCount64Subtract(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, subtract.OnesCount64)
}

func BenchmarkOnesCount64CallGCC(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, gccbuiltin.OnesCount64)
}

func BenchmarkOnesCount64Pop3(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, onescount.OnesCount64Pop3)
}

func BenchmarkOnesCount64Pop5(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, onescount.OnesCount64Pop5)
}

func BenchmarkOnesCount64Hakmem(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, onescount.OnesCount64Hakmem)
}

func BenchmarkOnesCount64Asm(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64Supplier.Next, onescount.OnesCount64Asm)
}
