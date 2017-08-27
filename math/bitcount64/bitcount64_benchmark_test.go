package bitcount64_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bitcount64"
)

// --- uint64 ---

// BenchmarkBitCount64Naive-8        	10000000	       180 ns/op
// BenchmarkBitCount64CallGCC-8      	20000000	        64.6 ns/op
// BenchmarkBitCount64Pop0-8         	200000000	         7.14 ns/op
// BenchmarkBitCount64Pop1-8         	200000000	         6.04 ns/op
// BenchmarkBitCount64Pop1Alt-8      	300000000	         5.31 ns/op <- Best
// BenchmarkBitCount64Pop3-8         	300000000	         5.56 ns/op
// BenchmarkBitCount64Pop4-8         	50000000	        32.2 ns/op
// BenchmarkBitCount64Pop5-8         	 5000000	       304 ns/op
// BenchmarkBitCount64Pop5a-8        	50000000	        34.3 ns/op
// BenchmarkBitCount64Pop6-8         	200000000	         7.66 ns/op
// BenchmarkBitCount64Hakmem-8       	200000000	         6.06 ns/op

func BenchmarkBitCount64Naive(b *testing.B) {
	benchmark.UintFuncUint64WithRandom(b, bitcount64.Naive)
}

func BenchmarkBitCount64CallGCC(b *testing.B) {
	benchmark.UintFuncUint64WithRandom(b, bitcount64.CallGCC)
}

func BenchmarkBitCount64Pop0(b *testing.B) {
	benchmark.UintFuncUint64WithRandom(b, bitcount64.Pop0)
}

func BenchmarkBitCount64Pop1(b *testing.B) {
	benchmark.UintFuncUint64WithRandom(b, bitcount64.Pop1)
}

func BenchmarkBitCount64Pop1Alt(b *testing.B) {
	benchmark.UintFuncUint64WithRandom(b, bitcount64.Pop1Alt)
}

func BenchmarkBitCount64Pop3(b *testing.B) {
	benchmark.UintFuncUint64WithRandom(b, bitcount64.Pop3)
}

func BenchmarkBitCount64Pop4(b *testing.B) {
	benchmark.UintFuncUint64WithRandom(b, bitcount64.Pop4)
}

func BenchmarkBitCount64Pop5(b *testing.B) {
	benchmark.UintFuncUint64WithRandom(b, bitcount64.Pop5)
}

func BenchmarkBitCount64Pop5a(b *testing.B) {
	benchmark.UintFuncUint64WithRandom(b, bitcount64.Pop5a)
}

func BenchmarkBitCount64Pop6(b *testing.B) {
	benchmark.UintFuncUint64WithRandom(b, bitcount64.Pop6)
}

func BenchmarkBitCount64Hakmem(b *testing.B) {
	benchmark.UintFuncUint64WithRandom(b, bitcount64.Hakmem)
}

func BenchmarkBitCount64Asm(b *testing.B) {
	benchmark.UintFuncUint64WithRandom(b, bitcount64.Asm)
}
