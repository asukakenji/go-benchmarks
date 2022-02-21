package combination_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/combination"
)

// goos: darwin
// goarch: amd64
// pkg: github.com/asukakenji/go-benchmarks/combination
// cpu: Intel(R) Xeon(R) W-2191B CPU @ 2.30GHz
// BenchmarkCombination0_4-36               6646375               179.0 ns/op
// BenchmarkCombination0_8-36                979368              1147 ns/op
// BenchmarkCombination0_16-36                 6682            174610 ns/op
// BenchmarkCombination1_4-36               5131611               230.9 ns/op
// BenchmarkCombination1_8-36                764448              1485 ns/op
// BenchmarkCombination1_16-36                 5386            222081 ns/op
// BenchmarkCombination1_24-36                  100          37095443 ns/op
// BenchmarkCombination2_4-36               5274069               224.0 ns/op
// BenchmarkCombination2_8-36                795069              1412 ns/op
// BenchmarkCombination2_16-36                 5755            209536 ns/op
// BenchmarkCombination2_24-36                  100          35673170 ns/op
// BenchmarkCombination2a_4-36              5215923               223.2 ns/op
// BenchmarkCombination2a_8-36               797730              1417 ns/op
// BenchmarkCombination2a_16-36                5499            208167 ns/op
// BenchmarkCombination2a_24-36                 100          34846380 ns/op
// BenchmarkCombination3_4-36               5497132               214.7 ns/op <- Best
// BenchmarkCombination3_8-36                835213              1338 ns/op   <- Best
// BenchmarkCombination3_16-36                 6132            193821 ns/op   <- Best
// BenchmarkCombination3_24-36                  100          32456427 ns/op   <- Best
// PASS
// ok      github.com/asukakenji/go-benchmarks/combination 32.864s

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

func BenchmarkCombination0_4(b *testing.B) {
	benchmarkCombination(b, Combination0[int], 4)
}

func BenchmarkCombination0_8(b *testing.B) {
	benchmarkCombination(b, Combination0[int], 8)
}

func BenchmarkCombination0_16(b *testing.B) {
	benchmarkCombination(b, Combination0[int], 16)
}

// func BenchmarkCombination0_24(b *testing.B) {
// 	benchmarkCombination(b, Combination0[int], 24)
// }

// ---

func BenchmarkCombination1_4(b *testing.B) {
	benchmarkCombination(b, combination.Combination1[int], 4)
}

func BenchmarkCombination1_8(b *testing.B) {
	benchmarkCombination(b, combination.Combination1[int], 8)
}

func BenchmarkCombination1_16(b *testing.B) {
	benchmarkCombination(b, combination.Combination1[int], 16)
}

func BenchmarkCombination1_24(b *testing.B) {
	benchmarkCombination(b, combination.Combination1[int], 24)
}

// ---

func BenchmarkCombination2_4(b *testing.B) {
	benchmarkCombination(b, combination.Combination2[int], 4)
}

func BenchmarkCombination2_8(b *testing.B) {
	benchmarkCombination(b, combination.Combination2[int], 8)
}

func BenchmarkCombination2_16(b *testing.B) {
	benchmarkCombination(b, combination.Combination2[int], 16)
}

func BenchmarkCombination2_24(b *testing.B) {
	benchmarkCombination(b, combination.Combination2[int], 24)
}

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

// ---

func BenchmarkCombination3_4(b *testing.B) {
	benchmarkCombination(b, combination.Combination3[int], 4)
}

func BenchmarkCombination3_8(b *testing.B) {
	benchmarkCombination(b, combination.Combination3[int], 8)
}

func BenchmarkCombination3_16(b *testing.B) {
	benchmarkCombination(b, combination.Combination3[int], 16)
}

func BenchmarkCombination3_24(b *testing.B) {
	benchmarkCombination(b, combination.Combination3[int], 24)
}
