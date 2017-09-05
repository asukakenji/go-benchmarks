package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/constantsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount8WorstCaseCalibrateSupplier-8      	2000000000	         1.85 ns/op
// BenchmarkOnesCount8WorstCaseCalibrateBenchmarker-8   	500000000	         3.17 ns/op
// BenchmarkOnesCount8WorstCaseNaive-8                  	200000000	         9.02 ns/op
// BenchmarkOnesCount8WorstCaseTable-8                  	500000000	         3.17 ns/op <- Best
// BenchmarkOnesCount8WorstCaseStdlib-8                 	500000000	         3.43 ns/op <- Best
// BenchmarkOnesCount8WorstCasePop0-8                   	300000000	         4.37 ns/op
// BenchmarkOnesCount8WorstCasePop1-8                   	500000000	         3.98 ns/op
// BenchmarkOnesCount8WorstCasePop1A-8                  	500000000	         3.97 ns/op
// BenchmarkOnesCount8WorstCaseReset-8                  	200000000	         9.04 ns/op
// BenchmarkOnesCount8WorstCaseSubtract-8               	200000000	         8.97 ns/op

var uint8SupplierWorstCase = constantsupplier.NewUint8(^uint8(0))

func BenchmarkOnesCount8WorstCaseCalibrateSupplier(b *testing.B) {
	benchmark.Uint8Supplier(b, uint8SupplierWorstCase.Next)
}

func BenchmarkOnesCount8WorstCaseCalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint8ToIntFunc(b, uint8SupplierWorstCase.Next)
}

func BenchmarkOnesCount8WorstCaseNaive(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, naive.OnesCount8)
}

func BenchmarkOnesCount8WorstCaseTable(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, table.OnesCount8)
}

func BenchmarkOnesCount8WorstCaseStdlib(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, stdlib.OnesCount8)
}

func BenchmarkOnesCount8WorstCasePop0(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop0.OnesCount8)
}

func BenchmarkOnesCount8WorstCasePop1(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop1.OnesCount8)
}

func BenchmarkOnesCount8WorstCasePop1A(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop1a.OnesCount8)
}

func BenchmarkOnesCount8WorstCaseReset(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, reset.OnesCount8)
}

func BenchmarkOnesCount8WorstCaseSubtract(b *testing.B) {
	uint8Supplier.Reset()
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, subtract.OnesCount8)
}
