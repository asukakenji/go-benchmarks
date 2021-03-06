package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/constantsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop9"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/table"
)

// BenchmarkOnesCount16WorstCaseCalibrateSupplier-8      	2000000000	         1.87 ns/op
// BenchmarkOnesCount16WorstCaseCalibrateBenchmarker-8   	1000000000	         2.92 ns/op
// BenchmarkOnesCount16WorstCaseNaive-8                  	100000000	        15.2 ns/op
// BenchmarkOnesCount16WorstCaseTable-8                  	500000000	         3.62 ns/op <- Best
// BenchmarkOnesCount16WorstCaseStdlib-8                 	500000000	         3.68 ns/op <- Best
// BenchmarkOnesCount16WorstCasePop0-8                   	300000000	         4.74 ns/op
// BenchmarkOnesCount16WorstCasePop1-8                   	300000000	         4.36 ns/op
// BenchmarkOnesCount16WorstCasePop1A-8                  	300000000	         4.38 ns/op
// BenchmarkOnesCount16WorstCaseReset-8                  	100000000	        10.9 ns/op
// BenchmarkOnesCount16WorstCaseSubtract-8               	100000000	        11.4 ns/op
// BenchmarkOnesCount15WorstCasePop9-8                   	500000000	         3.54 ns/op <- Best (15)
// BenchmarkOnesCount15WorstCasePop9Unrolled-8           	500000000	         3.54 ns/op <- Best (15)
// BenchmarkOnesCount16WorstCasePop9-8                   	500000000	         3.89 ns/op
// BenchmarkOnesCount16WorstCasePop9Unrolled-8           	500000000	         3.89 ns/op

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

func BenchmarkOnesCount15WorstCasePop9(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16SupplierWorstCase.Next, pop9.OnesCount15)
}

func BenchmarkOnesCount15WorstCasePop9Unrolled(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16SupplierWorstCase.Next, pop9.OnesCount15Unrolled)
}

func BenchmarkOnesCount16WorstCasePop9(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16SupplierWorstCase.Next, pop9.OnesCount16)
}

func BenchmarkOnesCount16WorstCasePop9Unrolled(b *testing.B) {
	benchmark.Uint16ToIntFunc(b, uint16SupplierWorstCase.Next, pop9.OnesCount16Unrolled)
}
