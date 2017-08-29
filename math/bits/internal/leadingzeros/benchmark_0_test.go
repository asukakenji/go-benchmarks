package leadingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/table"
)

// BenchmarkLeadingZeros0Calibrate-8   	1000000000	         2.38 ns/op
// BenchmarkLeadingZeros0Naive-8       	50000000	        36.8 ns/op
// BenchmarkLeadingZeros0Table-8       	300000000	         5.41 ns/op <- Best
// BenchmarkLeadingZeros0Stdlib-8      	300000000	         5.61 ns/op

func BenchmarkLeadingZeros0Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUintWithRandom(b)
}

func BenchmarkLeadingZeros0Naive(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, naive.LeadingZeros)
}

func BenchmarkLeadingZeros0Table(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, table.LeadingZerosConcept)
}

func BenchmarkLeadingZeros0Stdlib(b *testing.B) {
	benchmark.IntFuncUintWithRandom(b, stdlib.LeadingZeros)
}
