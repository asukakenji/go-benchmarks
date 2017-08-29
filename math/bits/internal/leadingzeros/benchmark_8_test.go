package leadingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/table"
)

// BenchmarkLeadingZeros8Calibrate-8   	1000000000	         2.21 ns/op
// BenchmarkLeadingZeros8Naive-8       	100000000	        10.0 ns/op
// BenchmarkLeadingZeros8Table-8       	300000000	         4.27 ns/op <- Best
// BenchmarkLeadingZeros8Stdlib-8      	300000000	         4.49 ns/op

func BenchmarkLeadingZeros8Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUintWithRandom(b)
}

func BenchmarkLeadingZeros8Naive(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, naive.LeadingZeros8)
}

func BenchmarkLeadingZeros8Table(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, table.LeadingZeros8)
}

func BenchmarkLeadingZeros8Stdlib(b *testing.B) {
	benchmark.IntFuncUint8WithRandom(b, stdlib.LeadingZeros8)
}
