package leadingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/table"
)

// BenchmarkLeadingZeros64Calibrate-8   	1000000000	         2.38 ns/op
// BenchmarkLeadingZeros64Naive-8       	300000000	         4.69 ns/op
// BenchmarkLeadingZeros64Table-8       	300000000	         4.47 ns/op <- Best
// BenchmarkLeadingZeros64Stdlib-8      	300000000	         5.47 ns/op

func BenchmarkLeadingZeros64Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUintWithRandom(b)
}

func BenchmarkLeadingZeros64Naive(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, naive.LeadingZeros64)
}

func BenchmarkLeadingZeros64Table(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, table.LeadingZeros64)
}

func BenchmarkLeadingZeros64Stdlib(b *testing.B) {
	benchmark.IntFuncUint64WithRandom(b, stdlib.LeadingZeros64)
}
