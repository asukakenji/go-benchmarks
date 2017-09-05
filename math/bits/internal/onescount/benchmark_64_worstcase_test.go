package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/constantsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/gccbuiltin"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/hakmem"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop3"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop5"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount64WorstCaseCalibrateSupplier-8      	2000000000	         1.86 ns/op
// BenchmarkOnesCount64WorstCaseCalibrateBenchmarker-8   	1000000000	         2.94 ns/op
// BenchmarkOnesCount64WorstCaseNaive-8                  	30000000	        46.8 ns/op
// BenchmarkOnesCount64WorstCaseTable-8                  	200000000	         6.62 ns/op
// BenchmarkOnesCount64WorstCaseStdlib-8                 	300000000	         5.14 ns/op
// BenchmarkOnesCount64WorstCasePop0-8                   	300000000	         5.95 ns/op
// BenchmarkOnesCount64WorstCasePop1-8                   	300000000	         4.92 ns/op
// BenchmarkOnesCount64WorstCasePop1Alt-8                	300000000	         4.38 ns/op <- Best
// BenchmarkOnesCount64WorstCaseReset-8                  	30000000	        50.6 ns/op
// BenchmarkOnesCount64WorstCaseSubtract-8               	30000000	        44.0 ns/op
// BenchmarkOnesCount64WorstCaseCallGCC-8                	20000000	        64.4 ns/op
// BenchmarkOnesCount64WorstCasePop3-8                   	300000000	         4.55 ns/op <- Best
// BenchmarkOnesCount64WorstCasePop5-8                   	 5000000	       309 ns/op
// BenchmarkOnesCount64WorstCaseHakmem-8                 	300000000	         5.15 ns/op
// BenchmarkOnesCount64WorstCaseAsm-8                    	300000000	         4.52 ns/op <- Best

var uint64SupplierWorstCase = constantsupplier.NewUint64(^uint64(0))

func BenchmarkOnesCount64WorstCaseCalibrateSupplier(b *testing.B) {
	benchmark.Uint64Supplier(b, uint64SupplierWorstCase.Next)
}

func BenchmarkOnesCount64WorstCaseCalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint64ToIntFunc(b, uint64SupplierWorstCase.Next)
}

func BenchmarkOnesCount64WorstCaseNaive(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, naive.OnesCount64)
}

func BenchmarkOnesCount64WorstCaseTable(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, table.OnesCount64)
}

func BenchmarkOnesCount64WorstCaseStdlib(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, stdlib.OnesCount64)
}

func BenchmarkOnesCount64WorstCasePop0(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, pop0.OnesCount64)
}

func BenchmarkOnesCount64WorstCasePop1(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, pop1.OnesCount64)
}

func BenchmarkOnesCount64WorstCasePop1Alt(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, pop1a.OnesCount64)
}

func BenchmarkOnesCount64WorstCaseReset(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, reset.OnesCount64)
}

func BenchmarkOnesCount64WorstCaseSubtract(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, subtract.OnesCount64)
}

func BenchmarkOnesCount64WorstCaseCallGCC(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, gccbuiltin.OnesCount64)
}

func BenchmarkOnesCount64WorstCasePop3(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, pop3.OnesCount64)
}

func BenchmarkOnesCount64WorstCasePop5(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, pop5.OnesCount64)
}

func BenchmarkOnesCount64WorstCaseHakmem(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, hakmem.OnesCount64)
}

func BenchmarkOnesCount64WorstCaseAsm(b *testing.B) {
	uint64Supplier.Reset()
	benchmark.Uint64ToIntFunc(b, uint64SupplierWorstCase.Next, onescount.OnesCount64Asm)
}
