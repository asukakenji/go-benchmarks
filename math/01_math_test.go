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

func TestBitCountUint(t *testing.T) {
	bitCountUintFuncs := []struct {
		name string
		f    func(uint) uint
	}{
		{"BitCountUintGCCImpl", math.BitCountUintGCCImpl},
		{"BitCountUintGCCImplSwitch", math.BitCountUintGCCImplSwitch},
	}
	gen := randomIntsInTheFirstPage()
	for _, bitCountUintFunc := range bitCountUintFuncs {
		for i := 0; i < 512; i++ {
			n := uint(gen())
			expected := math.BitCountUintNaive(n)
			got := bitCountUintFunc.f(n)
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUintFunc.name, n, got, expected)
			}
		}
	}
}

func TestBitCountUint32(t *testing.T) {
	bitCountUint32Funcs := []struct {
		name string
		f    func(uint32) uint
	}{
		{"BitCountUint32CallGCC", math.BitCountUint32CallGCC},
		{"BitCountUint32Pop0", math.BitCountUint32Pop0},
		{"BitCountUint32Pop1", math.BitCountUint32Pop1},
		{"BitCountUint32Pop1Alt", math.BitCountUint32Pop1Alt},
		{"BitCountUint32Pop2", math.BitCountUint32Pop2},
		{"BitCountUint32Pop2Alt", math.BitCountUint32Pop2Alt},
		{"BitCountUint32Pop3", math.BitCountUint32Pop3},
		{"BitCountUint32Pop4", math.BitCountUint32Pop4},
		{"BitCountUint32Pop5", math.BitCountUint32Pop5},
		{"BitCountUint32Pop5a", math.BitCountUint32Pop5a},
		{"BitCountUint32Pop6", math.BitCountUint32Pop6},
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
		{"BitCountUint64CallGCC", math.BitCountUint64CallGCC},
		{"BitCountUint64Pop1", math.BitCountUint64Pop1},
		{"BitCountUint64Pop1Alt", math.BitCountUint64Pop1Alt},
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

// --- uint ---

// BenchmarkBitCountUintNaive-8                      	10000000	       186 ns/op
// BenchmarkBitCountUintGCCImpl-8                    	100000000	        12.2 ns/op
// BenchmarkBitCountUintGCCImplSwitch-8              	100000000	        11.9 ns/op

func benchmarkBitCountUint(b *testing.B, f func(uint) uint) {
	gen := randomIntsInTheFirstPage()
	for i, count := 0, b.N; i < count; i++ {
		f(uint(gen()))
	}
}

func BenchmarkBitCountUintNaive(b *testing.B) {
	benchmarkBitCountUint(b, math.BitCountUintNaive)
}

func BenchmarkBitCountUintGCCImpl(b *testing.B) {
	benchmarkBitCountUint(b, math.BitCountUintGCCImpl)
}

func BenchmarkBitCountUintGCCImplSwitch(b *testing.B) {
	benchmarkBitCountUint(b, math.BitCountUintGCCImplSwitch)
}

// --- uint32 ---

// BenchmarkBitCountUint32Naive-8                    	20000000	        98.2 ns/op
// BenchmarkBitCountUint32CallGCC-8                  	20000000	        67.1 ns/op
// BenchmarkBitCountUint32Pop0-8                     	100000000	        12.6 ns/op
// BenchmarkBitCountUint32Pop1-8                     	100000000	        11.9 ns/op
// BenchmarkBitCountUint32Pop1Alt-8                  	100000000	        11.7 ns/op
// BenchmarkBitCountUint32Pop2-8                     	100000000	        12.3 ns/op
// BenchmarkBitCountUint32Pop2Alt-8                  	100000000	        11.9 ns/op
// BenchmarkBitCountUint32Pop3-8                     	100000000	        11.8 ns/op
// BenchmarkBitCountUint32Pop4-8                     	50000000	        27.4 ns/op
// BenchmarkBitCountUint32Pop5-8                     	10000000	       155 ns/op
// BenchmarkBitCountUint32Pop5a-8                    	50000000	        28.8 ns/op
// BenchmarkBitCountUint32Pop6-8                     	100000000	        11.7 ns/op

func benchmarkBitCountUint32(b *testing.B, f func(uint32) uint) {
	gen := randomIntsInTheFirstPage()
	for i, count := 0, b.N; i < count; i++ {
		f(uint32(gen()))
	}
}

func BenchmarkBitCountUint32Naive(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Naive)
}

func BenchmarkBitCountUint32CallGCC(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32CallGCC)
}

func BenchmarkBitCountUint32Pop0(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Pop0)
}

func BenchmarkBitCountUint32Pop1(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Pop1)
}

func BenchmarkBitCountUint32Pop1Alt(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Pop1Alt)
}

func BenchmarkBitCountUint32Pop2(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Pop2)
}

func BenchmarkBitCountUint32Pop2Alt(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Pop2Alt)
}

func BenchmarkBitCountUint32Pop3(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Pop3)
}

func BenchmarkBitCountUint32Pop4(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Pop4)
}

func BenchmarkBitCountUint32Pop5(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Pop5)
}

func BenchmarkBitCountUint32Pop5a(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Pop5a)
}

func BenchmarkBitCountUint32Pop6(b *testing.B) {
	benchmarkBitCountUint32(b, math.BitCountUint32Pop6)
}

// --- uint64 ---

// BenchmarkBitCountUint64Naive-8                    	10000000	       181 ns/op
// BenchmarkBitCountUint64CallGCC-8                  	20000000	        66.0 ns/op
// BenchmarkBitCountUint64Pop1-8                     	100000000	        12.6 ns/op
// BenchmarkBitCountUint64Pop1Alt-8                  	100000000	        12.2 ns/op

func benchmarkBitCountUint64(b *testing.B, f func(uint64) uint) {
	gen := randomIntsInTheFirstPage()
	for i, count := 0, b.N; i < count; i++ {
		f(uint64(gen()))
	}
}

func BenchmarkBitCountUint64Naive(b *testing.B) {
	benchmarkBitCountUint64(b, math.BitCountUint64Naive)
}

func BenchmarkBitCountUint64CallGCC(b *testing.B) {
	benchmarkBitCountUint64(b, math.BitCountUint64CallGCC)
}

func BenchmarkBitCountUint64Pop1(b *testing.B) {
	benchmarkBitCountUint64(b, math.BitCountUint64Pop1)
}

func BenchmarkBitCountUint64Pop1Alt(b *testing.B) {
	benchmarkBitCountUint64(b, math.BitCountUint64Pop1Alt)
}
