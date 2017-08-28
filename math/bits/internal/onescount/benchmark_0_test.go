package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/forward"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount0Calibrate-8         	1000000000	         2.36 ns/op
// BenchmarkOnesCount0Naive-8             	10000000	       183 ns/op
// BenchmarkOnesCount0Table-8             	200000000	         9.28 ns/op
// BenchmarkOnesCount0Pop1A-8             	200000000	         6.20 ns/op
// BenchmarkOnesCount0Reset-8             	50000000	        32.6 ns/op
// BenchmarkOnesCount0Subtract-8          	50000000	        35.5 ns/op
// BenchmarkOnesCount0IfConstBool-8       	300000000	         5.62 ns/op <- Best
// BenchmarkOnesCount0IfConstUint-8       	300000000	         5.63 ns/op
// BenchmarkOnesCount0SwitchConstUint-8   	300000000	         5.68 ns/op
// BenchmarkOnesCount0FuncPointer-8       	200000000	         7.42 ns/op

func BenchmarkOnesCount0Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUintWithRandom(b)
}

func BenchmarkOnesCount0Naive(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, naive.OnesCount)
}

func BenchmarkOnesCount0Table(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, table.OnesCountConcept)
}

func BenchmarkOnesCount0Pop1A(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, pop1a.OnesCount)
}

func BenchmarkOnesCount0Reset(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, reset.OnesCount)
}

func BenchmarkOnesCount0Subtract(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, subtract.OnesCount)
}

func BenchmarkOnesCount0IfConstBool(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, forward.OnesCountIfConstBool)
}

func BenchmarkOnesCount0IfConstUint(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, forward.OnesCountIfConstUint)
}

func BenchmarkOnesCount0SwitchConstUint(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, forward.OnesCountSwitchConstUint)
}

func BenchmarkOnesCount0FuncPointer(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, forward.OnesCountFuncPointer)
}
