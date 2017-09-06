package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop7"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop8"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop9"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount8CalibrateSupplier-8               	500000000	         4.03 ns/op
// BenchmarkOnesCount8CalibrateBenchmarker-8            	300000000	         5.00 ns/op
// BenchmarkOnesCount8Naive-8                           	50000000	        30.8 ns/op
// BenchmarkOnesCount8Table-8                           	300000000	         5.40 ns/op <- Best
// BenchmarkOnesCount8Stdlib-8                          	300000000	         5.38 ns/op <- Best
// BenchmarkOnesCount8Pop0-8                            	200000000	         7.02 ns/op
// BenchmarkOnesCount8Pop1-8                            	200000000	         6.44 ns/op
// BenchmarkOnesCount8Pop1A-8                           	200000000	         6.44 ns/op
// BenchmarkOnesCount8Reset-8                           	100000000	        15.5 ns/op
// BenchmarkOnesCount8Subtract-8                        	100000000	        13.7 ns/op
// BenchmarkOnesCount8Pop7-8                            	300000000	         5.86 ns/op
// BenchmarkOnesCount8Pop7Unrolled-8                    	300000000	         5.83 ns/op
// BenchmarkOnesCount7Pop8-8                            	300000000	         5.61 ns/op <- Best (7)
// BenchmarkOnesCount7Pop8Unrolled-8                    	300000000	         5.64 ns/op <- Best (7)
// BenchmarkOnesCount8Pop8-8                            	200000000	         6.25 ns/op
// BenchmarkOnesCount8Pop8Unrolled-8                    	200000000	         6.24 ns/op
// BenchmarkOnesCount8Pop9-8                            	300000000	         5.65 ns/op
// BenchmarkOnesCount8Pop9Unrolled-8                    	300000000	         5.71 ns/op

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

func BenchmarkOnesCount8Pop7(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, pop7.OnesCount8)
}

func BenchmarkOnesCount8Pop7Unrolled(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, pop7.OnesCount8Unrolled)
}

func BenchmarkOnesCount7Pop8(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, pop8.OnesCount7)
}

func BenchmarkOnesCount7Pop8Unrolled(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, pop8.OnesCount7Unrolled)
}

func BenchmarkOnesCount8Pop8(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, pop8.OnesCount8)
}

func BenchmarkOnesCount8Pop8Unrolled(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, pop8.OnesCount8Unrolled)
}

func BenchmarkOnesCount8Pop9(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, pop9.OnesCount8)
}

func BenchmarkOnesCount8Pop9Unrolled(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8Supplier.Next, pop9.OnesCount8Unrolled)
}
