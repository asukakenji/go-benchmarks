package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/forward"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount0Calibrate-8         	1000000000	         2.56 ns/op
// BenchmarkOnesCount0Naive-8             	10000000	       192 ns/op
// BenchmarkOnesCount0Table-8             	200000000	         9.35 ns/op
// BenchmarkOnesCount0Stdlib-8            	200000000	         6.74 ns/op
// BenchmarkOnesCount0Pop1A-8             	200000000	         6.41 ns/op
// BenchmarkOnesCount0Reset-8             	50000000	        31.9 ns/op
// BenchmarkOnesCount0Subtract-8          	30000000	        44.7 ns/op
// BenchmarkOnesCount0IfConstBool-8       	300000000	         5.75 ns/op <- Best
// BenchmarkOnesCount0IfConstUint-8       	300000000	         5.78 ns/op <- Best
// BenchmarkOnesCount0SwitchConstUint-8   	300000000	         5.75 ns/op <- Best
// BenchmarkOnesCount0FuncPointer-8       	200000000	         7.71 ns/op

var uintSupplier = randomsupplier.NewUint()

func BenchmarkLeadingZeros0CalibrateSupplier(b *testing.B) {
	benchmark.UintSupplier(b, uintSupplier.Next)
}

func BenchmarkLeadingZeros0CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUintToIntFunc(b, uintSupplier.Next)
}

func BenchmarkOnesCount0Naive(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, naive.OnesCount)
}

func BenchmarkOnesCount0Table(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, table.OnesCountConcept)
}

func BenchmarkOnesCount0Stdlib(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, stdlib.OnesCount)
}

func BenchmarkOnesCount0Pop1A(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, pop1a.OnesCount)
}

func BenchmarkOnesCount0Reset(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, reset.OnesCount)
}

func BenchmarkOnesCount0Subtract(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, subtract.OnesCount)
}

func BenchmarkOnesCount0IfConstBool(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, forward.OnesCountIfConstBool)
}

func BenchmarkOnesCount0IfConstUint(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, forward.OnesCountIfConstUint)
}

func BenchmarkOnesCount0SwitchConstUint(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, forward.OnesCountSwitchConstUint)
}

func BenchmarkOnesCount0FuncPointer(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, forward.OnesCountFuncPointer)
}
