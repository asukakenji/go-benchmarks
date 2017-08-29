package leadingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/table"
)

// BenchmarkLeadingZeros32Calibrate-8   	1000000000	         2.43 ns/op
// BenchmarkLeadingZeros32Naive-8       	300000000	         5.10 ns/op
// BenchmarkLeadingZeros32Table-8       	300000000	         4.57 ns/op <- Best
// BenchmarkLeadingZeros32Stdlib-8      	300000000	         5.40 ns/op

func BenchmarkLeadingZeros32Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUintWithRandom(b)
}

func BenchmarkLeadingZeros32Naive(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, naive.LeadingZeros32)
}

func BenchmarkLeadingZeros32Table(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, table.LeadingZeros32)
}

func BenchmarkLeadingZeros32Stdlib(b *testing.B) {
	benchmark.IntFuncUint32WithRandom(b, stdlib.LeadingZeros32)
}
