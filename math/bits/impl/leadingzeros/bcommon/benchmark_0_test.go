package bcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz3"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/table"
)

// goos: darwin
// goarch: amd64
// pkg: github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/bcommon
// cpu: Intel(R) Xeon(R) W-2191B CPU @ 2.30GHz
// BenchmarkLeadingZeros0Naive-36          91035858                11.78 ns/op
// BenchmarkLeadingZeros0Table-36          197893312                6.066 ns/op
// BenchmarkLeadingZeros0Stdlib-36         454678651                2.456 ns/op <- Best
// BenchmarkLeadingZeros0Nlz2a-36          165568584                7.286 ns/op
// BenchmarkLeadingZeros0Nlz3-36           71783720                15.57 ns/op
// PASS
// ok      github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/bcommon 7.614s

func BenchmarkLeadingZeros0Naive(b *testing.B) {
	benchmarkLeadingZeros(b, naive.LeadingZeros)
}

func BenchmarkLeadingZeros0Table(b *testing.B) {
	benchmarkLeadingZeros(b, table.LeadingZeros)
}

func BenchmarkLeadingZeros0Stdlib(b *testing.B) {
	benchmarkLeadingZeros(b, stdlib.LeadingZeros)
}

func BenchmarkLeadingZeros0Nlz2a(b *testing.B) {
	benchmarkLeadingZeros(b, nlz2a.LeadingZeros)
}

func BenchmarkLeadingZeros0Nlz3(b *testing.B) {
	benchmarkLeadingZeros(b, nlz3.LeadingZeros)
}
