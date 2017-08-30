package math_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math"
)

// BenchmarkIsOddNaive-8              	300000000	         4.47 ns/op
// BenchmarkIsOdd0-8                  	300000000	         4.19 ns/op <- Best
// BenchmarkIsOdd1-8                  	300000000	         4.20 ns/op

func BenchmarkIsOddNaive(b *testing.B) {
	intSupplier.Reset()
	benchmark.IntPredicate(b, intSupplier.Next, math.IsOddNaive)
}

func BenchmarkIsOdd0(b *testing.B) {
	intSupplier.Reset()
	benchmark.IntPredicate(b, intSupplier.Next, math.IsOdd0)
}

func BenchmarkIsOdd1(b *testing.B) {
	intSupplier.Reset()
	benchmark.IntPredicate(b, intSupplier.Next, math.IsOdd1)
}

// BenchmarkIsEvenNaive-8             	300000000	         4.48 ns/op
// BenchmarkIsEven0-8                 	300000000	         4.14 ns/op <- Best
// BenchmarkIsEven1-8                 	300000000	         4.19 ns/op

func BenchmarkIsEvenNaive(b *testing.B) {
	intSupplier.Reset()
	benchmark.IntPredicate(b, intSupplier.Next, math.IsEvenNaive)
}

func BenchmarkIsEven0(b *testing.B) {
	intSupplier.Reset()
	benchmark.IntPredicate(b, intSupplier.Next, math.IsEven0)
}

func BenchmarkIsEven1(b *testing.B) {
	intSupplier.Reset()
	benchmark.IntPredicate(b, intSupplier.Next, math.IsEven1)
}
