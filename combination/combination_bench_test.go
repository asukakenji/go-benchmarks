package combination_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/combination"
)

// goos: darwin
// goarch: amd64
// pkg: github.com/asukakenji/go-benchmarks/combination
// cpu: Intel(R) Xeon(R) W-2191B CPU @ 2.30GHz
// BenchmarkCombination1_4-36               5131611               230.9 ns/op
// BenchmarkCombination1_8-36                764448              1485 ns/op
// BenchmarkCombination1_16-36                 5386            222081 ns/op
// BenchmarkCombination1_24-36                  100          37095443 ns/op
// BenchmarkCombination2_4-36               5274069               224.0 ns/op <- Best (considering style)
// BenchmarkCombination2_8-36                795069              1412 ns/op   <- Best (considering style)
// BenchmarkCombination2_16-36                 5755            209536 ns/op   <- Best (considering style)
// BenchmarkCombination2_24-36                  100          35673170 ns/op   <- Best (considering style)
// BenchmarkCombination2a_4-36              5215923               223.2 ns/op
// BenchmarkCombination2a_8-36               797730              1417 ns/op
// BenchmarkCombination2a_16-36                5499            208167 ns/op
// BenchmarkCombination2a_24-36                 100          34846380 ns/op
// PASS
// ok      github.com/asukakenji/go-benchmarks/combination 30.600s

// Exported (global) variable serving as input for some
// of the benchmarks to ensure side-effect free calls
// are not optimized away.
var Input int

// Exported (global) variable to store function results
// during benchmarking to ensure side-effect free calls
// are not optimized away.
var Output int

func benchmarkCombination(b *testing.B, combination func([]int, int) [][]int, n int) {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	var s int
	for i := 0; i < b.N; i++ {
		s += combination(a, (i*97)%(n-1)+1)[0][0]
	}
	Output = s
}

// ===

// func BenchmarkCombination1_4(b *testing.B) {
// 	benchmarkCombination(b, combination.Combination1[int], 4)
// }

// func BenchmarkCombination1_8(b *testing.B) {
// 	benchmarkCombination(b, combination.Combination1[int], 8)
// }

// func BenchmarkCombination1_16(b *testing.B) {
// 	benchmarkCombination(b, combination.Combination1[int], 16)
// }

// func BenchmarkCombination1_24(b *testing.B) {
// 	benchmarkCombination(b, combination.Combination1[int], 24)
// }

// ---

// func BenchmarkCombination2_4(b *testing.B) {
// 	benchmarkCombination(b, combination.Combination2[int], 4)
// }

// func BenchmarkCombination2_8(b *testing.B) {
// 	benchmarkCombination(b, combination.Combination2[int], 8)
// }

// func BenchmarkCombination2_16(b *testing.B) {
// 	benchmarkCombination(b, combination.Combination2[int], 16)
// }

// func BenchmarkCombination2_24(b *testing.B) {
// 	benchmarkCombination(b, combination.Combination2[int], 24)
// }

// ---

func BenchmarkCombination2a_4(b *testing.B) {
	benchmarkCombination(b, combination.Combination2a[int], 4)
}

func BenchmarkCombination2a_8(b *testing.B) {
	benchmarkCombination(b, combination.Combination2a[int], 8)
}

func BenchmarkCombination2a_16(b *testing.B) {
	benchmarkCombination(b, combination.Combination2a[int], 16)
}

func BenchmarkCombination2a_24(b *testing.B) {
	benchmarkCombination(b, combination.Combination2a[int], 24)
}
