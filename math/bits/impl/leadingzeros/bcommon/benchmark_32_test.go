package bcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz1"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz10"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz10b"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz1a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz3"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz4"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz5"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/table"
)

// goos: darwin
// goarch: amd64
// pkg: github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/bcommon
// cpu: Intel(R) Xeon(R) W-2191B CPU @ 2.30GHz
// BenchmarkLeadingZeros32Naive-36                 137775746                8.811 ns/op
// BenchmarkLeadingZeros32Table-36                 586769715                2.041 ns/op <- Best
// BenchmarkLeadingZeros32Stdlib-36                532300576                2.250 ns/op <- Best
// BenchmarkLeadingZeros32Nlz1-36                  299750893                3.927 ns/op
// BenchmarkLeadingZeros32Nlz1a-36                 329702836                3.635 ns/op
// BenchmarkLeadingZeros32Nlz2-36                  385146025                3.093 ns/op <- Best (Non-Table)
// BenchmarkLeadingZeros32Nlz2a-36                 144831112                8.266 ns/op
// BenchmarkLeadingZeros32Nlz3-36                  143492344                8.227 ns/op
// BenchmarkLeadingZeros32Nlz4-36                  282813709                4.248 ns/op
// BenchmarkLeadingZeros32Nlz5-36                  381663351                3.129 ns/op <- Best (Non-Table)
// BenchmarkLeadingZeros32Nlz10-36                 532006430                2.234 ns/op <- Best
// BenchmarkLeadingZeros32Nlz10NoMultiply-36       534420279                2.235 ns/op <- Best
// BenchmarkLeadingZeros32Nlz10b-36                519854356                2.288 ns/op <- Best
// BenchmarkLeadingZeros32Nlz10bNoMultiply-36      420457243                2.843 ns/op <- Best
// PASS
// ok      github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/bcommon 22.806s

func BenchmarkLeadingZeros32Naive(b *testing.B) {
	benchmarkLeadingZeros32(b, naive.LeadingZeros32)
}

func BenchmarkLeadingZeros32Table(b *testing.B) {
	benchmarkLeadingZeros32(b, table.LeadingZeros32)
}

func BenchmarkLeadingZeros32Stdlib(b *testing.B) {
	benchmarkLeadingZeros32(b, stdlib.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz1(b *testing.B) {
	benchmarkLeadingZeros32(b, nlz1.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz1a(b *testing.B) {
	benchmarkLeadingZeros32(b, nlz1a.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz2(b *testing.B) {
	benchmarkLeadingZeros32(b, nlz2.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz2a(b *testing.B) {
	benchmarkLeadingZeros32(b, nlz2a.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz3(b *testing.B) {
	benchmarkLeadingZeros32(b, nlz3.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz4(b *testing.B) {
	benchmarkLeadingZeros32(b, nlz4.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz5(b *testing.B) {
	benchmarkLeadingZeros32(b, nlz5.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz10(b *testing.B) {
	benchmarkLeadingZeros32(b, nlz10.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz10NoMultiply(b *testing.B) {
	benchmarkLeadingZeros32(b, nlz10.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz10b(b *testing.B) {
	benchmarkLeadingZeros32(b, nlz10b.LeadingZeros32)
}

func BenchmarkLeadingZeros32Nlz10bNoMultiply(b *testing.B) {
	benchmarkLeadingZeros32(b, nlz10b.LeadingZeros32NoMultiply)
}
