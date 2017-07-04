package math_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math"
)

func TestBitCountUintNaive(t *testing.T) {
	cases := []struct {
		n        uint
		expected uint
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 1},
		{9, 2},
		{10, 2},
		{11, 3},
		{12, 2},
		{13, 3},
		{14, 3},
		{15, 4},
		{16, 1},
		{0x11, 2},
		{0x22, 2},
		{0x44, 2},
		{0x88, 2},
		{0xff, 8},
		{0x11111111, 8},
		{0x22222222, 8},
		{0x44444444, 8},
		{0x88888888, 8},
		{0xffffffff, 32},
		{0x12345678, 13},
	}
	for _, c := range cases {
		got := math.BitCountUintNaive(c.n)
		if got != c.expected {
			t.Errorf("BitCountUintNaive(%d) = %d, expected %d", c.n, got, c.expected)
		}
	}
}

func TestBitCountUint32Naive(t *testing.T) {
	cases := []struct {
		n        uint32
		expected uint
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 1},
		{9, 2},
		{10, 2},
		{11, 3},
		{12, 2},
		{13, 3},
		{14, 3},
		{15, 4},
		{16, 1},
		{0x11, 2},
		{0x22, 2},
		{0x44, 2},
		{0x88, 2},
		{0xff, 8},
		{0x11111111, 8},
		{0x22222222, 8},
		{0x44444444, 8},
		{0x88888888, 8},
		{0xffffffff, 32},
		{0x12345678, 13},
	}
	for _, c := range cases {
		got := math.BitCountUint32Naive(c.n)
		if got != c.expected {
			t.Errorf("BitCountUint32Naive(%d) = %d, expected %d", c.n, got, c.expected)
		}
	}
}

func TestBitCountUint64Naive(t *testing.T) {
	cases := []struct {
		n        uint64
		expected uint
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 1},
		{9, 2},
		{10, 2},
		{11, 3},
		{12, 2},
		{13, 3},
		{14, 3},
		{15, 4},
		{16, 1},
		{0x11, 2},
		{0x22, 2},
		{0x44, 2},
		{0x88, 2},
		{0xff, 8},
		{0x11111111, 8},
		{0x22222222, 8},
		{0x44444444, 8},
		{0x88888888, 8},
		{0xffffffff, 32},
		{0x12345678, 13},
		{0x123456789abcdef, 32},
	}
	for _, c := range cases {
		got := math.BitCountUint64Naive(c.n)
		if got != c.expected {
			t.Errorf("BitCountUint64Naive(%d) = %d, expected %d", c.n, got, c.expected)
		}
	}
}

func TestBitCountUint32(t *testing.T) {
	bitCountUint32Funcs := []struct {
		name string
		f    func(uint32) uint
	}{
		{"BitCountUint32Java", math.BitCountUint32Java},
		{"BitCountUint32CallGCC", math.BitCountUint32CallGCC},
		{"BitCountUint32GCCImpl", math.BitCountUint32GCCImpl},
	}
	gen := randomIntsInTheFirstPage()
	for _, bitCountUint32Func := range bitCountUint32Funcs {
		for i := 0; i < 512; i++ {
			n := uint32(gen())
			expected := math.BitCountUint32Naive(n)
			got := bitCountUint32Func.f(n)
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUint32Func.name, n, got, expected)
			}
		}
	}
}

func TestBitCountUint64(t *testing.T) {
	bitCountUint64Funcs := []struct {
		name string
		f    func(uint64) uint
	}{
		{"BitCountUint64Java", math.BitCountUint64Java},
		{"BitCountUint64CallGCC", math.BitCountUint64CallGCC},
		{"BitCountUint64GCCImpl", math.BitCountUint64GCCImpl},
	}
	gen := randomIntsInTheFirstPage()
	for _, bitCountUint64Func := range bitCountUint64Funcs {
		for i := 0; i < 512; i++ {
			n := uint64(gen())
			expected := math.BitCountUint64Naive(n)
			got := bitCountUint64Func.f(n)
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUint64Func.name, n, got, expected)
			}
		}
	}
}

func benchmarkBitCountUint32(b *testing.B, f func(uint32) uint) {
	gen := randomIntsInTheFirstPage()
	for i, count := 0, b.N; i < count; i++ {
		f(uint32(gen()))
	}
}

func benchmarkBitCountUint64(b *testing.B, f func(uint64) uint) {
	gen := randomIntsInTheFirstPage()
	for i, count := 0, b.N; i < count; i++ {
		f(uint64(gen()))
	}
}

func BenchmarkBitCountUint32Naive(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Naive)
}

func BenchmarkBitCountUint64Naive(b *testing.B) {
	benchmarkBitCountUint64(b, math.BitCountUint64Naive)
}

// java.lang.Integer#bitCount
func BenchmarkBitCountUint32Java(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Java)
}

// java.lang.Long#bitCount
func BenchmarkBitCountUint64Java(b *testing.B) {
	benchmarkBitCountUint64(b, math.BitCountUint64Java)
}

func BenchmarkBitCountUint32CallGCC(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32CallGCC)
}

func BenchmarkBitCountUint64CallGCC(b *testing.B) {
	benchmarkBitCountUint64(b, math.BitCountUint64CallGCC)
}

func BenchmarkBitCountUint32GCCImpl(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32GCCImpl)
}

func BenchmarkBitCountUint64GCCImpl(b *testing.B) {
	benchmarkBitCountUint64(b, math.BitCountUint64GCCImpl)
}
