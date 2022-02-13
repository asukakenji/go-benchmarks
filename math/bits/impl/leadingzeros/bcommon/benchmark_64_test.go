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
// BenchmarkLeadingZeros64Naive-36         60385888                19.08 ns/op
// BenchmarkLeadingZeros64Table-36         403680799                2.968 ns/op <- Best
// BenchmarkLeadingZeros64Stdlib-36        475226946                2.472 ns/op <- Best
// BenchmarkLeadingZeros64Nlz1-36          264188792                4.554 ns/op
// BenchmarkLeadingZeros64Nlz1a-36         305881768                3.908 ns/op
// BenchmarkLeadingZeros64Nlz2-36          319975749                3.738 ns/op <- Best (Non-Table)
// BenchmarkLeadingZeros64Nlz2a-36         164567097                7.278 ns/op
// BenchmarkLeadingZeros64Nlz3-36          67108095                16.51 ns/op
// BenchmarkLeadingZeros64Nlz5-36          316458127                3.776 ns/op <- Best (Non-Table)
// PASS
// ok      github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/bcommon 13.838s

func BenchmarkLeadingZeros64Naive(b *testing.B) {
	benchmarkLeadingZeros64(b, naive.LeadingZeros64)
}

func BenchmarkLeadingZeros64Table(b *testing.B) {
	benchmarkLeadingZeros64(b, table.LeadingZeros64)
}

func BenchmarkLeadingZeros64Stdlib(b *testing.B) {
	benchmarkLeadingZeros64(b, stdlib.LeadingZeros64)
}

func BenchmarkLeadingZeros64Nlz1(b *testing.B) {
	benchmarkLeadingZeros64(b, nlz1.LeadingZeros64)
}

func BenchmarkLeadingZeros64Nlz1a(b *testing.B) {
	benchmarkLeadingZeros64(b, nlz1a.LeadingZeros64)
}

func BenchmarkLeadingZeros64Nlz2(b *testing.B) {
	benchmarkLeadingZeros64(b, nlz2.LeadingZeros64)
}

func BenchmarkLeadingZeros64Nlz2a(b *testing.B) {
	benchmarkLeadingZeros64(b, nlz2a.LeadingZeros64)
}

func BenchmarkLeadingZeros64Nlz3(b *testing.B) {
	benchmarkLeadingZeros64(b, nlz3.LeadingZeros64)
}

func BenchmarkLeadingZeros64Nlz5(b *testing.B) {
	benchmarkLeadingZeros64(b, nlz5.LeadingZeros64)
}
