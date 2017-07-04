package math_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math"
)

var table = [...]uint{
	0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,

	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,

	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,

	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
}

func TestBitCountUintNaive(t *testing.T) {
	for x, expected := range table {
		got := math.BitCountUintNaive(uint(x))
		if got != expected {
			t.Errorf("BitCountUintNaive(%d) = %d, expected %d", x, got, expected)
		}
	}
	cases := []struct {
		x        uint
		expected uint
	}{
		{0x11111111, 8},
		{0x22222222, 8},
		{0x44444444, 8},
		{0x88888888, 8},
		{0xffffffff, 32},
		{0x12345678, 13},
	}
	for _, c := range cases {
		got := math.BitCountUintNaive(c.x)
		if got != c.expected {
			t.Errorf("BitCountUintNaive(%d) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestBitCountUint32Naive(t *testing.T) {
	for x, expected := range table {
		got := math.BitCountUint32Naive(uint32(x))
		if got != expected {
			t.Errorf("BitCountUint32Naive(%d) = %d, expected %d", x, got, expected)
		}
	}
	cases := []struct {
		x        uint32
		expected uint
	}{
		{0x11111111, 8},
		{0x22222222, 8},
		{0x44444444, 8},
		{0x88888888, 8},
		{0xffffffff, 32},
		{0x12345678, 13},
	}
	for _, c := range cases {
		got := math.BitCountUint32Naive(c.x)
		if got != c.expected {
			t.Errorf("BitCountUint32Naive(%d) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestBitCountUint64Naive(t *testing.T) {
	for x, expected := range table {
		got := math.BitCountUint64Naive(uint64(x))
		if got != expected {
			t.Errorf("BitCountUint64Naive(%d) = %d, expected %d", x, got, expected)
		}
	}
	cases := []struct {
		x        uint64
		expected uint
	}{
		{0x11111111, 8},
		{0x22222222, 8},
		{0x44444444, 8},
		{0x88888888, 8},
		{0xffffffff, 32},
		{0x12345678, 13},
		{0x123456789abcdef, 32},
	}
	for _, c := range cases {
		got := math.BitCountUint64Naive(c.x)
		if got != c.expected {
			t.Errorf("BitCountUint64Naive(%d) = %d, expected %d", c.x, got, c.expected)
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
		for x, expected := range table {
			got := bitCountUintFunc.f(uint(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUintFunc.name, x, got, expected)
			}
		}
		for i := 0; i < 512; i++ {
			x := uint(gen())
			expected := math.BitCountUintNaive(x)
			got := bitCountUintFunc.f(x)
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUintFunc.name, x, got, expected)
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
		for x, expected := range table {
			got := bitCountUint32Func.f(uint32(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUint32Func.name, x, got, expected)
			}
		}
		for i := 0; i < 512; i++ {
			x := uint32(gen())
			expected := math.BitCountUint32Naive(x)
			got := bitCountUint32Func.f(x)
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUint32Func.name, x, got, expected)
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
		for x, expected := range table {
			got := bitCountUint64Func.f(uint64(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUint64Func.name, x, got, expected)
			}
		}
		for i := 0; i < 512; i++ {
			x := uint64(gen())
			expected := math.BitCountUint64Naive(x)
			got := bitCountUint64Func.f(x)
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUint64Func.name, x, got, expected)
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
