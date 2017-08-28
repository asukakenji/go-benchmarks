package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount8Calibrate-8   	1000000000	         2.25 ns/op
// BenchmarkOnesCount8Naive-8       	50000000	        28.4 ns/op
// BenchmarkOnesCount8Table-8       	500000000	         3.88 ns/op <- Best
// BenchmarkOnesCount8Pop0-8        	300000000	         5.22 ns/op
// BenchmarkOnesCount8Pop1-8        	300000000	         4.73 ns/op
// BenchmarkOnesCount8Pop1A-8       	300000000	         4.78 ns/op
// BenchmarkOnesCount8Reset-8       	100000000	        13.0 ns/op
// BenchmarkOnesCount8Subtract-8    	100000000	        12.3 ns/op

func BenchmarkOnesCount8Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUint8WithRandom(b)
}

func BenchmarkOnesCount8Naive(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, naive.OnesCount8)
}

func BenchmarkOnesCount8Table(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, table.OnesCount8)
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
