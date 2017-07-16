package bitcount32_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bitcount32"
)

// --- uint32 ---

// BenchmarkBitCount32Naive-8            	20000000	        92.5 ns/op
// BenchmarkBitCount32CallGCC-8          	20000000	        64.4 ns/op
// BenchmarkBitCount32Pop0-8             	200000000	         6.06 ns/op
// BenchmarkBitCount32Pop1-8             	300000000	         5.41 ns/op
// BenchmarkBitCount32Pop1Alt-8          	300000000	         4.91 ns/op <- Best
// BenchmarkBitCount32Pop2-8             	200000000	         6.02 ns/op
// BenchmarkBitCount32Pop2Alt-8          	300000000	         5.22 ns/op
// BenchmarkBitCount32Pop3-8             	300000000	         5.19 ns/op
// BenchmarkBitCount32Pop4-8             	100000000	        20.7 ns/op
// BenchmarkBitCount32Pop5-8             	10000000	       148 ns/op
// BenchmarkBitCount32Pop5a-8            	100000000	        22.7 ns/op
// BenchmarkBitCount32Pop6-8             	300000000	         5.14 ns/op
// BenchmarkBitCount32Hakmem-8           	200000000	         6.35 ns/op
// BenchmarkBitCount32HakmemUnrolled-8   	200000000	         6.41 ns/op

func BenchmarkBitCount32Naive(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.Naive)
}

func BenchmarkBitCount32CallGCC(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.CallGCC)
}

func BenchmarkBitCount32Pop0(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.Pop0)
}

func BenchmarkBitCount32Pop1(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.Pop1)
}

func BenchmarkBitCount32Pop1Alt(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.Pop1Alt)
}

func BenchmarkBitCount32Pop2(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.Pop2)
}

func BenchmarkBitCount32Pop2Alt(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.Pop2Alt)
}

func BenchmarkBitCount32Pop3(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.Pop3)
}

func BenchmarkBitCount32Pop4(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.Pop4)
}

func BenchmarkBitCount32Pop5(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.Pop5)
}

func BenchmarkBitCount32Pop5a(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.Pop5a)
}

func BenchmarkBitCount32Pop6(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.Pop6)
}

func BenchmarkBitCount32Hakmem(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.Hakmem)
}

func BenchmarkBitCount32HakmemUnrolled(b *testing.B) {
	benchmark.UintFuncUint32WithRandom(b, bitcount32.HakmemUnrolled)
}
