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
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

// BenchmarkOnesCount64Calibrate-8   	1000000000	         2.34 ns/op
// BenchmarkOnesCount64Naive-8       	10000000	       185 ns/op
// BenchmarkOnesCount64Table-8       	200000000	         7.89 ns/op
// BenchmarkOnesCount64Stdlib-8      	200000000	         6.42 ns/op
// BenchmarkOnesCount64Pop0-8        	200000000	         7.11 ns/op
// BenchmarkOnesCount64Pop1-8        	200000000	         6.21 ns/op
// BenchmarkOnesCount64Pop1Alt-8     	300000000	         5.37 ns/op <- Best
// BenchmarkOnesCount64Reset-8       	50000000	        32.4 ns/op
// BenchmarkOnesCount64Subtract-8    	30000000	        43.0 ns/op
// BenchmarkOnesCount64CallGCC-8     	20000000	        63.3 ns/op
// BenchmarkOnesCount64Pop3-8        	300000000	         5.58 ns/op
// BenchmarkOnesCount64Pop5-8        	 5000000	       308 ns/op
// BenchmarkOnesCount64Hakmem-8      	200000000	         6.12 ns/op
// BenchmarkOnesCount64Asm-8         	300000000	         5.50 ns/op

func BenchmarkOnesCount64Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUint64WithRandom(b)
}

func BenchmarkOnesCount64Naive(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, naive.OnesCount64)
}

func BenchmarkOnesCount64Table(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, table.OnesCount64)
}

func BenchmarkOnesCount64Stdlib(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, stdlib.OnesCount64)
}

func BenchmarkOnesCount64Pop0(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, pop0.OnesCount64)
}

func BenchmarkOnesCount64Pop1(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, pop1.OnesCount64)
}

func BenchmarkOnesCount64Pop1Alt(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, pop1a.OnesCount64)
}

func BenchmarkOnesCount64Reset(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, reset.OnesCount64)
}

func BenchmarkOnesCount64Subtract(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, subtract.OnesCount64)
}

func BenchmarkOnesCount64CallGCC(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, onescount.OnesCount64CallGCC)
}

func BenchmarkOnesCount64Pop3(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, onescount.OnesCount64Pop3)
}

func BenchmarkOnesCount64Pop5(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, onescount.OnesCount64Pop5)
}

func BenchmarkOnesCount64Hakmem(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, onescount.OnesCount64Hakmem)
}

func BenchmarkOnesCount64Asm(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, onescount.OnesCount64Asm)
}
