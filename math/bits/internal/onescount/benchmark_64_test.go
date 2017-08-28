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

// BenchmarkOnesCount64Calibrate-8   	1000000000	         2.09 ns/op
// BenchmarkOnesCount64Naive-8       	10000000	       175 ns/op
// BenchmarkOnesCount64Table-8       	200000000	         7.62 ns/op
// BenchmarkOnesCount64Pop0-8        	200000000	         7.01 ns/op
// BenchmarkOnesCount64Pop1-8        	200000000	         6.07 ns/op
// BenchmarkOnesCount64Pop1Alt-8     	300000000	         5.32 ns/op <- Best
// BenchmarkOnesCount64Reset-8       	50000000	        33.4 ns/op
// BenchmarkOnesCount64Subtract-8    	50000000	        35.7 ns/op
// BenchmarkOnesCount64CallGCC-8     	20000000	        64.9 ns/op
// BenchmarkOnesCount64Pop3-8        	300000000	         5.54 ns/op
// BenchmarkOnesCount64Pop5-8        	 5000000	       306 ns/op
// BenchmarkOnesCount64Hakmem-8      	200000000	         6.11 ns/op
// BenchmarkOnesCount64Asm-8         	300000000	         5.68 ns/op

func BenchmarkOnesCount64Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUint64WithRandom(b)
}

func BenchmarkOnesCount64Naive(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, naive.OnesCount64)
}

func BenchmarkOnesCount64Table(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, table.OnesCount64)
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
