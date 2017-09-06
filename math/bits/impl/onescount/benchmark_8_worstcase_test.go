package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/constantsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop7"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop8"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop9"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/table"
)

// BenchmarkOnesCount8WorstCaseCalibrateSupplier-8      	2000000000	         1.86 ns/op
// BenchmarkOnesCount8WorstCaseCalibrateBenchmarker-8   	500000000	         3.23 ns/op
// BenchmarkOnesCount8WorstCaseNaive-8                  	200000000	         9.02 ns/op
// BenchmarkOnesCount8WorstCaseTable-8                  	500000000	         3.18 ns/op <- Best
// BenchmarkOnesCount8WorstCaseStdlib-8                 	500000000	         3.21 ns/op <- Best
// BenchmarkOnesCount8WorstCasePop0-8                   	300000000	         4.39 ns/op
// BenchmarkOnesCount8WorstCasePop1-8                   	300000000	         4.02 ns/op
// BenchmarkOnesCount8WorstCasePop1A-8                  	500000000	         4.03 ns/op
// BenchmarkOnesCount8WorstCaseReset-8                  	200000000	         6.95 ns/op
// BenchmarkOnesCount8WorstCaseSubtract-8               	200000000	         7.02 ns/op
// BenchmarkOnesCount8WorstCasePop7-8                   	500000000	         3.54 ns/op
// BenchmarkOnesCount8WorstCasePop7Unrolled-8           	500000000	         3.56 ns/op
// BenchmarkOnesCount7WorstCasePop8-8                   	500000000	         3.47 ns/op <- Best (7)
// BenchmarkOnesCount7WorstCasePop8Unrolled-8           	500000000	         3.46 ns/op <- Best (7)
// BenchmarkOnesCount8WorstCasePop8-8                   	500000000	         3.88 ns/op
// BenchmarkOnesCount8WorstCasePop8Unrolled-8           	500000000	         3.84 ns/op
// BenchmarkOnesCount8WorstCasePop9-8                   	500000000	         3.74 ns/op
// BenchmarkOnesCount8WorstCasePop9Unrolled-8           	500000000	         3.74 ns/op

var uint8SupplierWorstCase = constantsupplier.NewUint8(^uint8(0))

func BenchmarkOnesCount8WorstCaseCalibrateSupplier(b *testing.B) {
	benchmark.Uint8Supplier(b, uint8SupplierWorstCase.Next)
}

func BenchmarkOnesCount8WorstCaseCalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUint8ToIntFunc(b, uint8SupplierWorstCase.Next)
}

func BenchmarkOnesCount8WorstCaseNaive(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, naive.OnesCount8)
}

func BenchmarkOnesCount8WorstCaseTable(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, table.OnesCount8)
}

func BenchmarkOnesCount8WorstCaseStdlib(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, stdlib.OnesCount8)
}

func BenchmarkOnesCount8WorstCasePop0(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop0.OnesCount8)
}

func BenchmarkOnesCount8WorstCasePop1(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop1.OnesCount8)
}

func BenchmarkOnesCount8WorstCasePop1A(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop1a.OnesCount8)
}

func BenchmarkOnesCount8WorstCaseReset(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, reset.OnesCount8)
}

func BenchmarkOnesCount8WorstCaseSubtract(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, subtract.OnesCount8)
}

func BenchmarkOnesCount8WorstCasePop7(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop7.OnesCount8)
}

func BenchmarkOnesCount8WorstCasePop7Unrolled(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop7.OnesCount8Unrolled)
}

func BenchmarkOnesCount7WorstCasePop8(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop8.OnesCount7)
}

func BenchmarkOnesCount7WorstCasePop8Unrolled(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop8.OnesCount7Unrolled)
}

func BenchmarkOnesCount8WorstCasePop8(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop8.OnesCount8)
}

func BenchmarkOnesCount8WorstCasePop8Unrolled(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop8.OnesCount8Unrolled)
}

func BenchmarkOnesCount8WorstCasePop9(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop9.OnesCount8)
}

func BenchmarkOnesCount8WorstCasePop9Unrolled(b *testing.B) {
	benchmark.Uint8ToIntFunc(b, uint8SupplierWorstCase.Next, pop9.OnesCount8Unrolled)
}
