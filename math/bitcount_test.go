package math_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/random"
	"github.com/asukakenji/go-benchmarks/math"
)

var byteToBitCountTable = [...]uint{
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

var naiveCases = [...]struct {
	x        uint32
	expected uint
}{
	{0x00000000, 0},
	{0x11111111, 8},
	{0x22222222, 8},
	{0x33333333, 16},
	{0x44444444, 8},
	{0x55555555, 16},
	{0x66666666, 16},
	{0x77777777, 24},
	{0x88888888, 8},
	{0x99999999, 16},
	{0xaaaaaaaa, 16},
	{0xbbbbbbbb, 24},
	{0xcccccccc, 16},
	{0xdddddddd, 24},
	{0xeeeeeeee, 24},
	{0xffffffff, 32},
	{0x01234567, 12},
	{0x76543210, 12},
	{0x89abcdef, 20},
	{0xfedcba98, 20},
}

var commonCases = [...]uint64{
	0x0000000000000000,
	0x1111111111111111,
	0x2222222222222222,
	0x3333333333333333,
	0x4444444444444444,
	0x5555555555555555,
	0x6666666666666666,
	0x7777777777777777,
	0x8888888888888888,
	0x9999999999999999,
	0xaaaaaaaaaaaaaaaa,
	0xbbbbbbbbbbbbbbbb,
	0xcccccccccccccccc,
	0xdddddddddddddddd,
	0xeeeeeeeeeeeeeeee,
	0xffffffffffffffff,
	0x0123456789abcdef,
	0xfedcba9876543210,
}

func TestBitCountUintNaive(t *testing.T) {
	for x, expected := range byteToBitCountTable {
		got := math.BitCountUintNaive(uint(x))
		if got != expected {
			t.Errorf("BitCountUintNaive(%d) = %d, expected %d", x, got, expected)
		}
	}
	for _, c := range naiveCases {
		got := math.BitCountUintNaive(uint(c.x))
		if got != c.expected {
			t.Errorf("BitCountUintNaive(%d) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestBitCountUint32Naive(t *testing.T) {
	for x, expected := range byteToBitCountTable {
		got := math.BitCountUint32Naive(uint32(x))
		if got != expected {
			t.Errorf("BitCountUint32Naive(%d) = %d, expected %d", x, got, expected)
		}
	}
	for _, c := range naiveCases {
		got := math.BitCountUint32Naive(uint32(c.x))
		if got != c.expected {
			t.Errorf("BitCountUint32Naive(%d) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestBitCountUint64Naive(t *testing.T) {
	for x, expected := range byteToBitCountTable {
		got := math.BitCountUint64Naive(uint64(x))
		if got != expected {
			t.Errorf("BitCountUint64Naive(%d) = %d, expected %d", x, got, expected)
		}
	}
	for _, c := range naiveCases {
		got := math.BitCountUint64Naive(uint64(c.x))
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
	for _, bitCountUintFunc := range bitCountUintFuncs {
		for x, expected := range byteToBitCountTable {
			got := bitCountUintFunc.f(uint(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUintFunc.name, x, got, expected)
			}
		}
		for _, x := range commonCases {
			expected := math.BitCountUintNaive(uint(x))
			got := bitCountUintFunc.f(uint(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUintFunc.name, x, got, expected)
			}
		}
		gen := random.NewUintGenerator()
		for i := 0; i < 512; i++ {
			x := gen.Next()
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
		{"BitCountUint32Hakmem", math.BitCountUint32Hakmem},
		{"BitCountUint32HakmemUnrolled", math.BitCountUint32HakmemUnrolled},
	}
	for _, bitCountUint32Func := range bitCountUint32Funcs {
		for x, expected := range byteToBitCountTable {
			got := bitCountUint32Func.f(uint32(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUint32Func.name, x, got, expected)
			}
		}
		for _, x := range commonCases {
			expected := math.BitCountUint32Naive(uint32(x))
			got := bitCountUint32Func.f(uint32(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUint32Func.name, x, got, expected)
			}
		}
		gen := random.NewUint32Generator()
		for i := 0; i < 512; i++ {
			x := gen.Next()
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
		{"BitCountUint64Pop0", math.BitCountUint64Pop0},
		{"BitCountUint64Pop1", math.BitCountUint64Pop1},
		{"BitCountUint64Pop1Alt", math.BitCountUint64Pop1Alt},
		{"BitCountUint64Pop3", math.BitCountUint64Pop3},
		{"BitCountUint64Pop4", math.BitCountUint64Pop4},
		{"BitCountUint64Pop5", math.BitCountUint64Pop5},
		{"BitCountUint64Pop5a", math.BitCountUint64Pop5a},
		{"BitCountUint64Pop6", math.BitCountUint64Pop6},
		{"BitCountUint64Hakmem", math.BitCountUint64Hakmem},
	}
	for _, bitCountUint64Func := range bitCountUint64Funcs {
		for x, expected := range byteToBitCountTable {
			got := bitCountUint64Func.f(uint64(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUint64Func.name, x, got, expected)
			}
		}
		for _, x := range commonCases {
			expected := math.BitCountUint64Naive(uint64(x))
			got := bitCountUint64Func.f(uint64(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUint64Func.name, x, got, expected)
			}
		}
		gen := random.NewUint64Generator()
		for i := 0; i < 512; i++ {
			x := gen.Next()
			expected := math.BitCountUint64Naive(x)
			got := bitCountUint64Func.f(x)
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", bitCountUint64Func.name, x, got, expected)
			}
		}
	}
}

// --- uint32 ---

// BenchmarkBitCountUint32Naive-8            	20000000	        94.6 ns/op
// BenchmarkBitCountUint32CallGCC-8          	20000000	        64.9 ns/op
// BenchmarkBitCountUint32Pop0-8             	200000000	         6.29 ns/op
// BenchmarkBitCountUint32Pop1-8             	300000000	         5.66 ns/op
// BenchmarkBitCountUint32Pop1Alt-8          	300000000	         5.79 ns/op
// BenchmarkBitCountUint32Pop2-8             	200000000	         6.26 ns/op
// BenchmarkBitCountUint32Pop2Alt-8          	300000000	         5.38 ns/op
// BenchmarkBitCountUint32Pop3-8             	300000000	         5.38 ns/op
// BenchmarkBitCountUint32Pop4-8             	100000000	        21.4 ns/op
// BenchmarkBitCountUint32Pop5-8             	10000000	       149 ns/op
// BenchmarkBitCountUint32Pop5a-8            	100000000	        22.3 ns/op
// BenchmarkBitCountUint32Pop6-8             	300000000	         5.39 ns/op
// BenchmarkBitCountUint32Hakmem-8           	200000000	         6.67 ns/op
// BenchmarkBitCountUint32HakmemUnrolled-8   	200000000	         6.71 ns/op

func BenchmarkBitCountUint32Naive(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32Naive)
}

func BenchmarkBitCountUint32CallGCC(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32CallGCC)
}

func BenchmarkBitCountUint32Pop0(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32Pop0)
}

func BenchmarkBitCountUint32Pop1(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32Pop1)
}

func BenchmarkBitCountUint32Pop1Alt(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32Pop1Alt)
}

func BenchmarkBitCountUint32Pop2(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32Pop2)
}

func BenchmarkBitCountUint32Pop2Alt(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32Pop2Alt)
}

func BenchmarkBitCountUint32Pop3(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32Pop3)
}

func BenchmarkBitCountUint32Pop4(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32Pop4)
}

func BenchmarkBitCountUint32Pop5(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32Pop5)
}

func BenchmarkBitCountUint32Pop5a(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32Pop5a)
}

func BenchmarkBitCountUint32Pop6(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32Pop6)
}

func BenchmarkBitCountUint32Hakmem(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32Hakmem)
}

func BenchmarkBitCountUint32HakmemUnrolled(b *testing.B) {
	benchmark.RandomUintFuncUint32(b, math.BitCountUint32HakmemUnrolled)
}

// --- uint64 ---

// BenchmarkBitCountUint64Naive-8            	10000000	       173 ns/op
// BenchmarkBitCountUint64CallGCC-8          	20000000	        65.3 ns/op
// BenchmarkBitCountUint64Pop1-8             	200000000	         6.07 ns/op
// BenchmarkBitCountUint64Pop1Alt-8          	300000000	         5.25 ns/op
// BenchmarkBitCountUint64Pop3-8             	300000000	         5.48 ns/op
// BenchmarkBitCountUint64Pop4-8             	50000000	        29.9 ns/op
// BenchmarkBitCountUint64Pop5-8             	 5000000	       301 ns/op
// BenchmarkBitCountUint64Pop5a-8            	50000000	        35.0 ns/op
// BenchmarkBitCountUint64Pop6-8             	200000000	         7.60 ns/op
// BenchmarkBitCountUint64Hakmem-8           	200000000	         6.04 ns/op

func BenchmarkBitCountUint64Naive(b *testing.B) {
	benchmark.RandomUintFuncUint64(b, math.BitCountUint64Naive)
}

func BenchmarkBitCountUint64CallGCC(b *testing.B) {
	benchmark.RandomUintFuncUint64(b, math.BitCountUint64CallGCC)
}

func BenchmarkBitCountUint64Pop1(b *testing.B) {
	benchmark.RandomUintFuncUint64(b, math.BitCountUint64Pop1)
}

func BenchmarkBitCountUint64Pop1Alt(b *testing.B) {
	benchmark.RandomUintFuncUint64(b, math.BitCountUint64Pop1Alt)
}

func BenchmarkBitCountUint64Pop3(b *testing.B) {
	benchmark.RandomUintFuncUint64(b, math.BitCountUint64Pop3)
}

func BenchmarkBitCountUint64Pop4(b *testing.B) {
	benchmark.RandomUintFuncUint64(b, math.BitCountUint64Pop4)
}

func BenchmarkBitCountUint64Pop5(b *testing.B) {
	benchmark.RandomUintFuncUint64(b, math.BitCountUint64Pop5)
}

func BenchmarkBitCountUint64Pop5a(b *testing.B) {
	benchmark.RandomUintFuncUint64(b, math.BitCountUint64Pop5a)
}

func BenchmarkBitCountUint64Pop6(b *testing.B) {
	benchmark.RandomUintFuncUint64(b, math.BitCountUint64Pop6)
}

func BenchmarkBitCountUint64Hakmem(b *testing.B) {
	benchmark.RandomUintFuncUint64(b, math.BitCountUint64Hakmem)
}

// --- uint ---

// BenchmarkBitCountUintNaive-8              	10000000	       179 ns/op
// BenchmarkBitCountUintGCCImpl-8            	200000000	         6.13 ns/op
// BenchmarkBitCountUintGCCImplSwitch-8      	300000000	         5.46 ns/op

func BenchmarkBitCountUintNaive(b *testing.B) {
	benchmark.RandomUintFuncUint(b, math.BitCountUintNaive)
}

func BenchmarkBitCountUintGCCImpl(b *testing.B) {
	benchmark.RandomUintFuncUint(b, math.BitCountUintGCCImpl)
}

func BenchmarkBitCountUintGCCImplSwitch(b *testing.B) {
	benchmark.RandomUintFuncUint(b, math.BitCountUintGCCImplSwitch)
}
