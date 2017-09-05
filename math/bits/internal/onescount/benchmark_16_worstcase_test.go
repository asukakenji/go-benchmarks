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

// BenchmarkOnesCount16WorstCaseCalibrateSupplier-8      	2000000000	         1.83 ns/op
// BenchmarkOnesCount16WorstCaseCalibrateBenchmarker-8   	1000000000	         2.87 ns/op
// BenchmarkOnesCount16WorstCaseNaive-8                  	100000000	        15.2 ns/op
// BenchmarkOnesCount16WorstCaseTable-8                  	500000000	         3.50 ns/op <- Best
// BenchmarkOnesCount16WorstCaseStdlib-8                 	500000000	         3.55 ns/op <- Best
// BenchmarkOnesCount16WorstCasePop0-8                   	300000000	         4.63 ns/op
// BenchmarkOnesCount16WorstCasePop1-8                   	300000000	         4.24 ns/op
// BenchmarkOnesCount16WorstCasePop1A-8                  	300000000	         4.27 ns/op
// BenchmarkOnesCount16WorstCaseReset-8                  	100000000	        10.7 ns/op
// BenchmarkOnesCount16WorstCaseSubtract-8               	100000000	        11.1 ns/op

var uint16SupplierWorstCase = constantsupplier.NewUint16(^uint16(0))

func BenchmarkOnesCount16WorstCaseCalibrateSupplier(b *testing.B) {
	benchmark.Uint16Supplier(b, uint16SupplierWorstCase.Next)
}

func BenchmarkOnesCount16WorstCaseCalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint16ToIntFunc(b, uint16SupplierWorstCase.Next)
}

func BenchmarkOnesCount16WorstCaseNaive(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16SupplierWorstCase.Next, naive.OnesCount16)
}

func BenchmarkOnesCount16WorstCaseTable(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16SupplierWorstCase.Next, table.OnesCount16)
}

func BenchmarkOnesCount16WorstCaseStdlib(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16SupplierWorstCase.Next, stdlib.OnesCount16)
}

func BenchmarkOnesCount16WorstCasePop0(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16SupplierWorstCase.Next, pop0.OnesCount16)
}

func BenchmarkOnesCount16WorstCasePop1(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16SupplierWorstCase.Next, pop1.OnesCount16)
}

func BenchmarkOnesCount16WorstCasePop1A(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16SupplierWorstCase.Next, pop1a.OnesCount16)
}

func BenchmarkOnesCount16WorstCaseReset(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16SupplierWorstCase.Next, reset.OnesCount16)
}

func BenchmarkOnesCount16WorstCaseSubtract(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16SupplierWorstCase.Next, subtract.OnesCount16)
}
