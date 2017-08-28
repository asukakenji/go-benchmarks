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

// BenchmarkOnesCount8Calibrate-8   	1000000000	         2.31 ns/op
// BenchmarkOnesCount8Naive-8       	50000000	        29.5 ns/op
// BenchmarkOnesCount8Table-8       	300000000	         4.13 ns/op <- Best
// BenchmarkOnesCount8Stdlib-8      	300000000	         4.13 ns/op <- Best
// BenchmarkOnesCount8Pop0-8        	300000000	         5.44 ns/op
// BenchmarkOnesCount8Pop1-8        	300000000	         5.01 ns/op
// BenchmarkOnesCount8Pop1A-8       	300000000	         4.92 ns/op
// BenchmarkOnesCount8Reset-8       	100000000	        13.8 ns/op
// BenchmarkOnesCount8Subtract-8    	100000000	        12.9 ns/op

func BenchmarkOnesCount8Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUint8WithRandom(b)
}

func BenchmarkOnesCount8Naive(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, naive.OnesCount8)
}

func BenchmarkOnesCount8Table(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, table.OnesCount8)
}

func BenchmarkOnesCount8Stdlib(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, stdlib.OnesCount8)
}

func BenchmarkOnesCount8Pop0(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, pop0.OnesCount8)
}

func BenchmarkOnesCount8Pop1(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, pop1.OnesCount8)
}

func BenchmarkOnesCount8Pop1A(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, pop1a.OnesCount8)
}

func BenchmarkOnesCount8Reset(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, reset.OnesCount8)
}

func BenchmarkOnesCount8Subtract(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, subtract.OnesCount8)
}
