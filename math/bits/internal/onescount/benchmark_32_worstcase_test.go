package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/constantsupplier"
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

// BenchmarkOnesCount32WorstCaseCalibrateSupplier-8      	2000000000	         1.83 ns/op
// BenchmarkOnesCount32WorstCaseCalibrateBenchmarker-8   	500000000	         3.13 ns/op
// BenchmarkOnesCount32WorstCaseNaive-8                  	100000000	        19.7 ns/op
// BenchmarkOnesCount32WorstCaseTable-8                  	300000000	         4.45 ns/op <- Best
// BenchmarkOnesCount32WorstCaseStdlib-8                 	300000000	         4.65 ns/op
// BenchmarkOnesCount32WorstCasePop0-8                   	300000000	         5.22 ns/op
// BenchmarkOnesCount32WorstCasePop1-8                   	300000000	         4.54 ns/op
// BenchmarkOnesCount32WorstCasePop1A-8                  	300000000	         4.20 ns/op <- Best
// BenchmarkOnesCount32WorstCaseReset-8                  	100000000	        19.3 ns/op
// BenchmarkOnesCount32WorstCaseSubtract-8               	100000000	        19.3 ns/op
// BenchmarkOnesCount32WorstCaseCallGCC-8                	20000000	        63.4 ns/op
// BenchmarkOnesCount32WorstCasePop2-8                   	300000000	         5.21 ns/op
// BenchmarkOnesCount32WorstCasePop2Alt-8                	300000000	         4.43 ns/op <- Best
// BenchmarkOnesCount32WorstCasePop3-8                   	300000000	         4.43 ns/op <- Best
// BenchmarkOnesCount32WorstCasePop5-8                   	10000000	       147 ns/op
// BenchmarkOnesCount32WorstCaseHakmem-8                 	300000000	         5.88 ns/op
// BenchmarkOnesCount32WorstCaseHakmemUnrolled-8         	300000000	         5.65 ns/op

var uint32SupplierWorstCase = constantsupplier.NewUint32(^uint32(0))

func BenchmarkOnesCount32WorstCaseCalibrateSupplier(b *testing.B) {
	benchmark.Uint32Supplier(b, uint32SupplierWorstCase.Next)
}

func BenchmarkOnesCount32WorstCaseCalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint32ToIntFunc(b, uint32SupplierWorstCase.Next)
}

func BenchmarkOnesCount32WorstCaseNaive(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, naive.OnesCount32)
}

func BenchmarkOnesCount32WorstCaseTable(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, table.OnesCount32)
}

func BenchmarkOnesCount32WorstCaseStdlib(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, stdlib.OnesCount32)
}

func BenchmarkOnesCount32WorstCasePop0(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, pop0.OnesCount32)
}

func BenchmarkOnesCount32WorstCasePop1(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, pop1.OnesCount32)
}

func BenchmarkOnesCount32WorstCasePop1A(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, pop1a.OnesCount32)
}

func BenchmarkOnesCount32WorstCaseReset(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, reset.OnesCount32)
}

func BenchmarkOnesCount32WorstCaseSubtract(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, subtract.OnesCount32)
}

func BenchmarkOnesCount32WorstCaseCallGCC(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, gccbuiltin.OnesCount32)
}

func BenchmarkOnesCount32WorstCasePop2(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, pop2.OnesCount32)
}

func BenchmarkOnesCount32WorstCasePop2Alt(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, pop2a.OnesCount32)
}

func BenchmarkOnesCount32WorstCasePop3(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, pop3.OnesCount32)
}

func BenchmarkOnesCount32WorstCasePop5(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, pop5.OnesCount32)
}

func BenchmarkOnesCount32WorstCaseHakmem(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, hakmem.OnesCount32)
}

func BenchmarkOnesCount32WorstCaseHakmemUnrolled(b *testing.B) {
	benchmark.Uint32ToIntFunc(b, uint32SupplierWorstCase.Next, hakmem.OnesCount32Unrolled)
}
