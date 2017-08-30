package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount32Calibrate-8        	1000000000	         2.09 ns/op
// BenchmarkOnesCount32Naive-8            	20000000	       100 ns/op
// BenchmarkOnesCount32Table-8            	300000000	         5.64 ns/op
// BenchmarkOnesCount32Stdlib-8           	300000000	         5.75 ns/op
// BenchmarkOnesCount32Pop0-8             	200000000	         6.40 ns/op
// BenchmarkOnesCount32Pop1-8             	300000000	         5.75 ns/op
// BenchmarkOnesCount32Pop1A-8            	300000000	         5.26 ns/op <- Best
// BenchmarkOnesCount32Reset-8            	100000000	        22.6 ns/op
// BenchmarkOnesCount32Subtract-8         	50000000	        32.5 ns/op
// BenchmarkOnesCount32CallGCC-8          	20000000	        67.3 ns/op
// BenchmarkOnesCount32Pop2-8             	200000000	         6.53 ns/op
// BenchmarkOnesCount32Pop2Alt-8          	300000000	         5.53 ns/op
// BenchmarkOnesCount32Pop3-8             	300000000	         5.56 ns/op
// BenchmarkOnesCount32Pop5-8             	10000000	       154 ns/op
// BenchmarkOnesCount32Hakmem-8           	200000000	         6.68 ns/op
// BenchmarkOnesCount32HakmemUnrolled-8   	200000000	         6.80 ns/op

var uint32Supplier = randomsupplier.NewUint32()

func BenchmarkLeadingZeros32CalibrateSupplier(b *testing.B) {
	benchmark.Uint32Supplier(b, uint32Supplier.Next)
}

func BenchmarkLeadingZeros32CalibrateBenchmarker(b *testing.B) {
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
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, onescount.OnesCount32CallGCC)
}

func BenchmarkOnesCount32Pop2(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, onescount.OnesCount32Pop2)
}

func BenchmarkOnesCount32Pop2Alt(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, onescount.OnesCount32Pop2Alt)
}

func BenchmarkOnesCount32Pop3(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, onescount.OnesCount32Pop3)
}

func BenchmarkOnesCount32Pop5(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, onescount.OnesCount32Pop5)
}

func BenchmarkOnesCount32Hakmem(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, onescount.OnesCount32Hakmem)
}

func BenchmarkOnesCount32HakmemUnrolled(b *testing.B) {
	uint32Supplier.Reset()
	benchmark.Uint32ToIntFunc(b, uint32Supplier.Next, onescount.OnesCount32HakmemUnrolled)
}
