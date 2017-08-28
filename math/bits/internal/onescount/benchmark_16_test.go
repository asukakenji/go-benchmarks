package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount16Calibrate-8   	1000000000	         2.32 ns/op
// BenchmarkOnesCount16Naive-8       	30000000	        55.6 ns/op
// BenchmarkOnesCount16Table-8       	300000000	         4.69 ns/op <- Best
// BenchmarkOnesCount16Stdlib-8      	300000000	         4.75 ns/op <- Best
// BenchmarkOnesCount16Pop0-8        	200000000	         5.92 ns/op
// BenchmarkOnesCount16Pop1-8        	300000000	         5.55 ns/op
// BenchmarkOnesCount16Pop1A-8       	300000000	         5.59 ns/op
// BenchmarkOnesCount16Reset-8       	100000000	        17.3 ns/op
// BenchmarkOnesCount16Subtract-8    	100000000	        17.2 ns/op

func BenchmarkOnesCount16Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUint16WithRandom(b)
}

func BenchmarkOnesCount16Naive(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, naive.OnesCount16)
}

func BenchmarkOnesCount16Table(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, table.OnesCount16)
}

func BenchmarkOnesCount16Stdlib(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, stdlib.OnesCount16)
}

func BenchmarkOnesCount16Pop0(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, pop0.OnesCount16)
}

func BenchmarkOnesCount16Pop1(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, pop1.OnesCount16)
}

func BenchmarkOnesCount16Pop1A(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, pop1a.OnesCount16)
}

func BenchmarkOnesCount16Reset(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, reset.OnesCount16)
}

func BenchmarkOnesCount16Subtract(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, subtract.OnesCount16)
}
