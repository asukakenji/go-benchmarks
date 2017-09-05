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

// BenchmarkOnesCount0CalibrateSupplier-8               	500000000	         3.64 ns/op
// BenchmarkOnesCount0CalibrateBenchmarker-8            	300000000	         4.82 ns/op
// BenchmarkOnesCount0Naive-8                           	10000000	       179 ns/op
// BenchmarkOnesCount0Table-8                           	100000000	        10.8 ns/op
// BenchmarkOnesCount0Stdlib-8                          	200000000	         7.63 ns/op
// BenchmarkOnesCount0Pop1A-8                           	200000000	         7.25 ns/op
// BenchmarkOnesCount0Reset-8                           	50000000	        34.5 ns/op
// BenchmarkOnesCount0Subtract-8                        	50000000	        37.2 ns/op
// BenchmarkOnesCount0IfConstBool-8                     	200000000	         6.70 ns/op <- Best
// BenchmarkOnesCount0IfConstUint-8                     	200000000	         6.68 ns/op <- Best
// BenchmarkOnesCount0SwitchConstUint-8                 	200000000	         6.67 ns/op <- Best
// BenchmarkOnesCount0FuncPointer-8                     	200000000	         8.95 ns/op

var uintSupplier = randomsupplier.NewUint()

func BenchmarkOnesCount0CalibrateSupplier(b *testing.B) {
	benchmark.UintSupplier(b, uintSupplier.Next)
}

func BenchmarkOnesCount0CalibrateBenchmarker(b *testing.B) {
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
