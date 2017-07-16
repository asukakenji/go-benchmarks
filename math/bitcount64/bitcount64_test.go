package bitcount64_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/random"
	"github.com/asukakenji/go-benchmarks/math/bitcount64"
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

func TestBitCount64Naive(t *testing.T) {
	cases := []struct {
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
	for x, expected := range byteToBitCountTable {
		got := bitcount64.Naive(uint64(x))
		if got != expected {
			t.Errorf("bitcount64.Naive(%d) = %d, expected %d", x, got, expected)
		}
	}
	for _, c := range cases {
		got := bitcount64.Naive(uint64(c.x))
		if got != c.expected {
			t.Errorf("bitcount64.Naive(%d) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestBitCount64(t *testing.T) {
	implementations := []struct {
		name string
		f    func(uint64) uint
	}{
		{"bitcount64.CallGCC", bitcount64.CallGCC},
		{"bitcount64.Pop0", bitcount64.Pop0},
		{"bitcount64.Pop1", bitcount64.Pop1},
		{"bitcount64.Pop1Alt", bitcount64.Pop1Alt},
		{"bitcount64.Pop3", bitcount64.Pop3},
		{"bitcount64.Pop4", bitcount64.Pop4},
		{"bitcount64.Pop5", bitcount64.Pop5},
		{"bitcount64.Pop5a", bitcount64.Pop5a},
		{"bitcount64.Pop6", bitcount64.Pop6},
		{"bitcount64.Hakmem", bitcount64.Hakmem},
	}
	cases := []uint64{
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
	for _, impl := range implementations {
		for x, expected := range byteToBitCountTable {
			got := impl.f(uint64(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", impl.name, x, got, expected)
			}
		}
		for _, x := range cases {
			expected := bitcount64.Naive(uint64(x))
			got := impl.f(uint64(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", impl.name, x, got, expected)
			}
		}
		gen := random.NewUint64Generator()
		for i := 0; i < 512; i++ {
			x := gen.Next()
			expected := bitcount64.Naive(x)
			got := impl.f(x)
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", impl.name, x, got, expected)
			}
		}
	}
}
