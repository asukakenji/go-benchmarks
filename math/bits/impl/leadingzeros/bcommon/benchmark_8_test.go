package bcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz1"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz1a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz3"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz5"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/table"
)

// goos: darwin
// goarch: amd64
// pkg: github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/bcommon
// cpu: Intel(R) Xeon(R) W-2191B CPU @ 2.30GHz
// BenchmarkLeadingZeros8Naive-36          375842836                3.180 ns/op
// BenchmarkLeadingZeros8Table-36          837485616                1.417 ns/op <- Best
// BenchmarkLeadingZeros8Stdlib-36         837043983                1.422 ns/op <- Best
// BenchmarkLeadingZeros8Nlz1-36           559964184                2.123 ns/op <- Best (Non-Table)
// BenchmarkLeadingZeros8Nlz1a-36          543472698                2.239 ns/op
// BenchmarkLeadingZeros8Nlz2-36           456479832                2.597 ns/op
// BenchmarkLeadingZeros8Nlz2a-36          298404591                3.987 ns/op
// BenchmarkLeadingZeros8Nlz3-36           513300402                2.328 ns/op
// BenchmarkLeadingZeros8Nlz5-36           559056831                2.129 ns/op <- Best (Non-Table)
// PASS
// ok      github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/bcommon 13.186s

func BenchmarkLeadingZeros8Naive(b *testing.B) {
	benchmarkLeadingZeros8(b, naive.LeadingZeros8)
}

func BenchmarkLeadingZeros8Table(b *testing.B) {
	benchmarkLeadingZeros8(b, table.LeadingZeros8)
}

func BenchmarkLeadingZeros8Stdlib(b *testing.B) {
	benchmarkLeadingZeros8(b, stdlib.LeadingZeros8)
}

func BenchmarkLeadingZeros8Nlz1(b *testing.B) {
	benchmarkLeadingZeros8(b, nlz1.LeadingZeros8)
}

func BenchmarkLeadingZeros8Nlz1a(b *testing.B) {
	benchmarkLeadingZeros8(b, nlz1a.LeadingZeros8)
}

func BenchmarkLeadingZeros8Nlz2(b *testing.B) {
	benchmarkLeadingZeros8(b, nlz2.LeadingZeros8)
}

func BenchmarkLeadingZeros8Nlz2a(b *testing.B) {
	benchmarkLeadingZeros8(b, nlz2a.LeadingZeros8)
}

func BenchmarkLeadingZeros8Nlz3(b *testing.B) {
	benchmarkLeadingZeros8(b, nlz3.LeadingZeros8)
}

func BenchmarkLeadingZeros8Nlz5(b *testing.B) {
	benchmarkLeadingZeros8(b, nlz5.LeadingZeros8)
}
