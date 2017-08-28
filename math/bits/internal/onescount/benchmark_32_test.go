package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop0"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/reset"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount32Calibrate-8        	1000000000	         2.04 ns/op
// BenchmarkOnesCount32Naive-8            	20000000	        94.1 ns/op
// BenchmarkOnesCount32Table-8            	300000000	         5.47 ns/op
// BenchmarkOnesCount32Pop0-8             	200000000	         6.30 ns/op
// BenchmarkOnesCount32Pop1-8             	300000000	         5.99 ns/op
// BenchmarkOnesCount32Pop1A-8            	300000000	         5.12 ns/op <- Best
// BenchmarkOnesCount32Reset-8            	100000000	        21.5 ns/op
// BenchmarkOnesCount32Subtract-8         	50000000	        31.5 ns/op
// BenchmarkOnesCount32CallGCC-8          	20000000	        64.7 ns/op
// BenchmarkOnesCount32Pop2-8             	200000000	         6.35 ns/op
// BenchmarkOnesCount32Pop2Alt-8          	300000000	         5.44 ns/op
// BenchmarkOnesCount32Pop3-8             	300000000	         5.40 ns/op
// BenchmarkOnesCount32Pop5-8             	10000000	       151 ns/op
// BenchmarkOnesCount32Hakmem-8           	200000000	         6.49 ns/op
// BenchmarkOnesCount32HakmemUnrolled-8   	200000000	         6.60 ns/op

func BenchmarkOnesCount32Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUint32WithRandom(b)
}

func BenchmarkOnesCount32Naive(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, naive.OnesCount32)
}

func BenchmarkOnesCount32Table(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, table.OnesCount32)
}

func BenchmarkOnesCount32Pop0(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, pop0.OnesCount32)
}

func BenchmarkOnesCount32Pop1(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, pop1.OnesCount32)
}

func BenchmarkOnesCount32Pop1A(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, pop1a.OnesCount32)
}

func BenchmarkOnesCount32Reset(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, reset.OnesCount32)
}

func BenchmarkOnesCount32Subtract(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, subtract.OnesCount32)
}

func BenchmarkOnesCount32CallGCC(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, onescount.OnesCount32CallGCC)
}

func BenchmarkOnesCount32Pop2(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, onescount.OnesCount32Pop2)
}

func BenchmarkOnesCount32Pop2Alt(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, onescount.OnesCount32Pop2Alt)
}

func BenchmarkOnesCount32Pop3(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, onescount.OnesCount32Pop3)
}

func BenchmarkOnesCount32Pop5(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, onescount.OnesCount32Pop5)
}

func BenchmarkOnesCount32Hakmem(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, onescount.OnesCount32Hakmem)
}

func BenchmarkOnesCount32HakmemUnrolled(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, onescount.OnesCount32HakmemUnrolled)
}
