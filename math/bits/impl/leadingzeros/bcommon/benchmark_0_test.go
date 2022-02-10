package bcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz3"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/table"
)

// BenchmarkLeadingZeros0CalibrateSupplier-8      	500000000	         3.95 ns/op
// BenchmarkLeadingZeros0CalibrateBenchmarker-8   	300000000	         4.91 ns/op
// BenchmarkLeadingZeros0Naive-8                  	50000000	        35.9 ns/op
// BenchmarkLeadingZeros0Table-8                  	200000000	         6.42 ns/op <- Best
// BenchmarkLeadingZeros0Stdlib-8                 	200000000	         6.94 ns/op <- Best
// BenchmarkLeadingZeros0Nlz2a-8                  	100000000	        13.4 ns/op
// BenchmarkLeadingZeros0Nlz3-8                   	50000000	        30.4 ns/op

var uintSupplier = randomsupplier.NewUint()

func BenchmarkLeadingZeros0CalibrateSupplier(b *testing.B) {
	benchmark.UintSupplier(b, uintSupplier.Next)
}

func BenchmarkLeadingZeros0CalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateUintToIntFunc(b, uintSupplier.Next)
}

func BenchmarkLeadingZeros0Naive(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, naive.LeadingZeros)
}

func BenchmarkLeadingZeros0Table(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, table.LeadingZeros)
}

func BenchmarkLeadingZeros0Stdlib(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, stdlib.LeadingZeros)
}

func BenchmarkLeadingZeros0Nlz2a(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, nlz2a.LeadingZeros)
}

func BenchmarkLeadingZeros0Nlz3(b *testing.B) {
	uintSupplier.Reset()
	benchmark.UintToIntFunc(b, uintSupplier.Next, nlz3.LeadingZeros)
}
