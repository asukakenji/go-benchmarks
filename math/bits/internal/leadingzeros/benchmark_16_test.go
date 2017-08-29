package leadingzeros_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/table"
)

// BenchmarkLeadingZeros16Calibrate-8   	1000000000	         2.36 ns/op
// BenchmarkLeadingZeros16Naive-8       	200000000	         9.63 ns/op
// BenchmarkLeadingZeros16Table-8       	300000000	         4.49 ns/op <- Best
// BenchmarkLeadingZeros16Stdlib-8      	300000000	         5.14 ns/op

func BenchmarkLeadingZeros16Calibrate(b *testing.B) {
	benchmark.CalibrateAnyFuncUintWithRandom(b)
}

func BenchmarkLeadingZeros16Naive(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, naive.LeadingZeros16)
}

func BenchmarkLeadingZeros16Table(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, table.LeadingZeros16)
}

func BenchmarkLeadingZeros16Stdlib(b *testing.B) {
	benchmark.IntFuncUint16WithRandom(b, stdlib.LeadingZeros16)
}
