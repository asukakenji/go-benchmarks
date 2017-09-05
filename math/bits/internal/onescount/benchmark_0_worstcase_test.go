package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/constantsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/forward"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount0WorstCaseCalibrateSupplier-8      	2000000000	         1.86 ns/op
// BenchmarkOnesCount0WorstCaseCalibrateBenchmarker-8   	1000000000	         2.92 ns/op
// BenchmarkOnesCount0WorstCaseNaive-8                  	30000000	        46.8 ns/op
// BenchmarkOnesCount0WorstCaseTable-8                  	200000000	         8.79 ns/op
// BenchmarkOnesCount0WorstCaseStdlib-8                 	300000000	         5.14 ns/op
// BenchmarkOnesCount0WorstCasePop1A-8                  	300000000	         4.93 ns/op
// BenchmarkOnesCount0WorstCaseReset-8                  	30000000	        46.7 ns/op
// BenchmarkOnesCount0WorstCaseSubtract-8               	30000000	        44.0 ns/op
// BenchmarkOnesCount0WorstCaseIfConstBool-8            	300000000	         4.37 ns/op <- Best
// BenchmarkOnesCount0WorstCaseIfConstUint-8            	300000000	         4.36 ns/op <- Best
// BenchmarkOnesCount0WorstCaseSwitchConstUint-8        	300000000	         4.36 ns/op <- Best
// BenchmarkOnesCount0WorstCaseFuncPointer-8            	200000000	         6.54 ns/op

var uintSupplierWorstCase = constantsupplier.NewUint(^uint(0))

func BenchmarkOnesCount0WorstCaseCalibrateSupplier(b *testing.B) {
	benchmark.UintSupplier(b, uintSupplierWorstCase.Next)
}

func BenchmarkOnesCount0WorstCaseCalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUintToIntFunc(b, uintSupplierWorstCase.Next)
}

func BenchmarkOnesCount0WorstCaseNaive(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplierWorstCase.Next, naive.OnesCount)
}

func BenchmarkOnesCount0WorstCaseTable(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplierWorstCase.Next, table.OnesCountConcept)
}

func BenchmarkOnesCount0WorstCaseStdlib(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplierWorstCase.Next, stdlib.OnesCount)
}

func BenchmarkOnesCount0WorstCasePop1A(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplierWorstCase.Next, pop1a.OnesCount)
}

func BenchmarkOnesCount0WorstCaseReset(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplierWorstCase.Next, reset.OnesCount)
}

func BenchmarkOnesCount0WorstCaseSubtract(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplierWorstCase.Next, subtract.OnesCount)
}

func BenchmarkOnesCount0WorstCaseIfConstBool(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplierWorstCase.Next, forward.OnesCountIfConstBool)
}

func BenchmarkOnesCount0WorstCaseIfConstUint(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplierWorstCase.Next, forward.OnesCountIfConstUint)
}

func BenchmarkOnesCount0WorstCaseSwitchConstUint(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplierWorstCase.Next, forward.OnesCountSwitchConstUint)
}

func BenchmarkOnesCount0WorstCaseFuncPointer(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplierWorstCase.Next, forward.OnesCountFuncPointer)
}
