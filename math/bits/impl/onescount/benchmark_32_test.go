package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/gccbuiltin"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/hakmem"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop2"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop2a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop3"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop5"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/table"
)

// BenchmarkOnesCount32CalibrateSupplier-8               	500000000	         3.25 ns/op
// BenchmarkOnesCount32CalibrateBenchmarker-8            	300000000	         4.39 ns/op
// BenchmarkOnesCount32Naive-8                           	20000000	        97.6 ns/op
// BenchmarkOnesCount32Table-8                           	200000000	         6.55 ns/op <- Best
// BenchmarkOnesCount32Stdlib-8                          	200000000	         6.83 ns/op
// BenchmarkOnesCount32Pop0-8                            	200000000	         7.62 ns/op
// BenchmarkOnesCount32Pop1-8                            	200000000	         6.77 ns/op
// BenchmarkOnesCount32Pop1A-8                           	200000000	         6.21 ns/op <- Best
// BenchmarkOnesCount32Reset-8                           	100000000	        23.8 ns/op
// BenchmarkOnesCount32Subtract-8                        	100000000	        23.9 ns/op
// BenchmarkOnesCount32CallGCC-8                         	20000000	        67.0 ns/op
// BenchmarkOnesCount32Pop2-8                            	200000000	         7.56 ns/op
// BenchmarkOnesCount32Pop2Alt-8                         	200000000	         6.54 ns/op <- Best
// BenchmarkOnesCount32Pop3-8                            	200000000	         6.53 ns/op <- Best
// BenchmarkOnesCount32Pop5-8                            	10000000	       146 ns/op
// BenchmarkOnesCount32Hakmem-8                          	200000000	         8.04 ns/op
// BenchmarkOnesCount32HakmemUnrolled-8                  	200000000	         7.90 ns/op

var uint32Supplier = randomsupplier.NewUint32()

func BenchmarkOnesCount32CalibrateSupplier(b *testing.B) {
	benchmark.Uint32Supplier(b, uint32Supplier.Next)
}

func BenchmarkOnesCount32CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint32ToIntFunc(b, uint32Supplier.Next)
}

func BenchmarkOnesCount32Naive(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, naive.OnesCount32)
}

func BenchmarkOnesCount32Table(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, table.OnesCount32)
}

func BenchmarkOnesCount32Stdlib(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, stdlib.OnesCount32)
}

func BenchmarkOnesCount32Pop0(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, pop0.OnesCount32)
}

func BenchmarkOnesCount32Pop1(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, pop1.OnesCount32)
}

func BenchmarkOnesCount32Pop1A(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, pop1a.OnesCount32)
}

func BenchmarkOnesCount32Reset(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, reset.OnesCount32)
}

func BenchmarkOnesCount32Subtract(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, subtract.OnesCount32)
}

func BenchmarkOnesCount32CallGCC(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, gccbuiltin.OnesCount32)
}

func BenchmarkOnesCount32Pop2(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, pop2.OnesCount32)
}

func BenchmarkOnesCount32Pop2Alt(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, pop2a.OnesCount32)
}

func BenchmarkOnesCount32Pop3(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, pop3.OnesCount32)
}

func BenchmarkOnesCount32Pop5(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, pop5.OnesCount32)
}

func BenchmarkOnesCount32Hakmem(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, hakmem.OnesCount32)
}

func BenchmarkOnesCount32HakmemUnrolled(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, hakmem.OnesCount32Unrolled)
}
