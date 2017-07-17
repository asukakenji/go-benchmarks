package permutation_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/permutation"
)

var (
	index   = uint(0)
	numbers = [][]uint{
		{14556, 3742, 9070, 26821, 11803, 31083, 28105, 7365, 32544},
		{31549},
		{1296, 28906, 9632, 27000, 26974, 13051, 4223, 7056},
		{25443, 4555},
		{17178, 14913, 3075, 11320, 4865, 22683, 21602},
		{30115, 6416, 1653},
		{15827, 32039, 14348, 9896, 10925, 10605},
		{22174, 31173, 8201, 31479},
		{27431, 3402, 11322, 20749, 18733},
	}
	n = uint(len(numbers))
)

func next() []uint {
	index++
	if index == n {
		index = 0
	}
	return numbers[index]
}

func BenchmarkUintSliceGenerator(b *testing.B) {
	index = 0
	benchmark.CalibrateUintSliceGenerator(b, next)
}

// BenchmarkPermutation0-8         	    3000	    485235 ns/op
// BenchmarkPermutation0Alt1-8     	    3000	    484863 ns/op
// BenchmarkPermutation1Alt1-8     	    3000	    440356 ns/op
// BenchmarkPermutation1Alt2-8     	    3000	    443558 ns/op
// BenchmarkPermutation1Alt3-8     	    3000	    439344 ns/op
// BenchmarkPermutation1Alt4-8     	    3000	    435427 ns/op <- Best
// BenchmarkPermutation1Alt5-8     	    3000	    446428 ns/op
// BenchmarkPermutation1Alt6-8     	    3000	    436928 ns/op

func BenchmarkPermutation0(b *testing.B) {
	index = 0
	benchmark.UintSliceGenerator(b, next, permutation.Permutation0)
}

func BenchmarkPermutation0Alt1(b *testing.B) {
	index = 0
	benchmark.UintSliceGenerator(b, next, permutation.Permutation0Alt1)
}

func BenchmarkPermutation1Alt1(b *testing.B) {
	index = 0
	benchmark.UintSliceGenerator(b, next, permutation.Permutation1Alt1)
}

func BenchmarkPermutation1Alt2(b *testing.B) {
	index = 0
	benchmark.UintSliceGenerator(b, next, permutation.Permutation1Alt2)
}

func BenchmarkPermutation1Alt3(b *testing.B) {
	index = 0
	benchmark.UintSliceGenerator(b, next, permutation.Permutation1Alt3)
}

func BenchmarkPermutation1Alt4(b *testing.B) {
	index = 0
	benchmark.UintSliceGenerator(b, next, permutation.Permutation1Alt4)
}

func BenchmarkPermutation1Alt5(b *testing.B) {
	index = 0
	benchmark.UintSliceGenerator(b, next, permutation.Permutation1Alt5)
}

func BenchmarkPermutation1Alt6(b *testing.B) {
	index = 0
	benchmark.UintSliceGenerator(b, next, permutation.Permutation1Alt6)
}
