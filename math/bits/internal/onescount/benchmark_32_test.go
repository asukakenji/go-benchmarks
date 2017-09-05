package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/gccbuiltin"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/hakmem"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop2"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop2a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop3"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop5"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount32CalibrateSupplier-8               	500000000	         3.27 ns/op
// BenchmarkOnesCount32CalibrateBenchmarker-8            	300000000	         4.73 ns/op
// BenchmarkOnesCount32Naive-8                           	20000000	        98.9 ns/op
// BenchmarkOnesCount32Table-8                           	200000000	         7.02 ns/op
// BenchmarkOnesCount32Stdlib-8                          	200000000	         7.18 ns/op
// BenchmarkOnesCount32Pop0-8                            	200000000	         8.03 ns/op
// BenchmarkOnesCount32Pop1-8                            	200000000	         7.22 ns/op
// BenchmarkOnesCount32Pop1A-8                           	200000000	         6.72 ns/op <- Best
// BenchmarkOnesCount32Reset-8                           	50000000	        23.7 ns/op
// BenchmarkOnesCount32Subtract-8                        	50000000	        30.6 ns/op
// BenchmarkOnesCount32CallGCC-8                         	20000000	        64.2 ns/op
// BenchmarkOnesCount32Pop2-8                            	200000000	         7.88 ns/op
// BenchmarkOnesCount32Pop2Alt-8                         	200000000	         7.05 ns/op
// BenchmarkOnesCount32Pop3-8                            	200000000	         6.98 ns/op <- Best
// BenchmarkOnesCount32Pop5-8                            	10000000	       155 ns/op
// BenchmarkOnesCount32Hakmem-8                          	200000000	         8.39 ns/op
// BenchmarkOnesCount32HakmemUnrolled-8                  	200000000	         8.24 ns/op

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
