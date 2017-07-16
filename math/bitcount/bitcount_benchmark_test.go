package bitcount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bitcount"
)

// --- uint ---

// BenchmarkBitCountNaive-8           	10000000	       182 ns/op
// BenchmarkBitCountPop1Alt-8         	200000000	         6.29 ns/op
// BenchmarkBitCountPop4-8            	50000000	        31.3 ns/op
// BenchmarkBitCountPop5a-8           	50000000	        34.4 ns/op
// BenchmarkBitCountPop1AltSwitch-8   	300000000	         5.54 ns/op <- Best

func BenchmarkBitCountNaive(b *testing.B) {
	benchmark.RandomUintFuncUint(b, bitcount.Naive)
}

func BenchmarkBitCountPop1Alt(b *testing.B) {
	benchmark.RandomUintFuncUint(b, bitcount.Pop1Alt)
}

func BenchmarkBitCountPop4(b *testing.B) {
	benchmark.RandomUintFuncUint(b, bitcount.Pop4)
}

func BenchmarkBitCountPop5a(b *testing.B) {
	benchmark.RandomUintFuncUint(b, bitcount.Pop5a)
}

func BenchmarkBitCountPop1AltSwitch(b *testing.B) {
	benchmark.RandomUintFuncUint(b, bitcount.Pop1AltSwitch)
}
