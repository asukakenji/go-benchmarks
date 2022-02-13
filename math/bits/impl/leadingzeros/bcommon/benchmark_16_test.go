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
// BenchmarkLeadingZeros16Naive-36         285502256                4.246 ns/op
// BenchmarkLeadingZeros16Table-36         634337701                1.882 ns/op <- Best
// BenchmarkLeadingZeros16Stdlib-36        545357505                2.184 ns/op <- Best
// BenchmarkLeadingZeros16Nlz1-36          362951762                3.294 ns/op
// BenchmarkLeadingZeros16Nlz1a-36         402672145                2.959 ns/op
// BenchmarkLeadingZeros16Nlz2-36          421391158                2.829 ns/op
// BenchmarkLeadingZeros16Nlz2a-36         215555361                5.521 ns/op
// BenchmarkLeadingZeros16Nlz3-36          298482657                4.003 ns/op
// BenchmarkLeadingZeros16Nlz5-36          452581203                2.645 ns/op <- Best (Non-Table)
// PASS
// ok      github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/bcommon 14.031s

func BenchmarkLeadingZeros16Naive(b *testing.B) {
	benchmarkLeadingZeros16(b, naive.LeadingZeros16)
}

func BenchmarkLeadingZeros16Table(b *testing.B) {
	benchmarkLeadingZeros16(b, table.LeadingZeros16)
}

func BenchmarkLeadingZeros16Stdlib(b *testing.B) {
	benchmarkLeadingZeros16(b, stdlib.LeadingZeros16)
}

func BenchmarkLeadingZeros16Nlz1(b *testing.B) {
	benchmarkLeadingZeros16(b, nlz1.LeadingZeros16)
}

func BenchmarkLeadingZeros16Nlz1a(b *testing.B) {
	benchmarkLeadingZeros16(b, nlz1a.LeadingZeros16)
}

func BenchmarkLeadingZeros16Nlz2(b *testing.B) {
	benchmarkLeadingZeros16(b, nlz2.LeadingZeros16)
}

func BenchmarkLeadingZeros16Nlz2a(b *testing.B) {
	benchmarkLeadingZeros16(b, nlz2a.LeadingZeros16)
}

func BenchmarkLeadingZeros16Nlz3(b *testing.B) {
	benchmarkLeadingZeros16(b, nlz3.LeadingZeros16)
}

func BenchmarkLeadingZeros16Nlz5(b *testing.B) {
	benchmarkLeadingZeros16(b, nlz5.LeadingZeros16)
}
