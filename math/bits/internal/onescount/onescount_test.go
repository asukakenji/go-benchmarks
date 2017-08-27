package onescount

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/random"
)

func TestOnesCountNaive(t *testing.T) {
	cases := []struct {
		x        uint32
		expected int
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
		got := OnesCountNaive(uint(x))
		if got != expected {
			t.Errorf("OnesCountNaive(%d) = %d, expected %d", x, got, expected)
		}
	}
	for _, c := range cases {
		got := OnesCountNaive(uint(c.x))
		if got != c.expected {
			t.Errorf("OnesCountNaive(%d) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestOnesCountAll(t *testing.T) {
	implementations := []struct {
		name string
		f    func(uint) int
	}{
		{"OnesCountPop1Alt", OnesCountPop1Alt},
		{"OnesCountPop4", OnesCountPop4},
		{"OnesCountPop5a", OnesCountPop5a},
		{"OnesCountPop1AltSwitch", OnesCountPop1AltSwitch},
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
			got := impl.f(uint(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", impl.name, x, got, expected)
			}
		}
		for _, x := range cases {
			expected := OnesCountNaive(uint(x))
			got := impl.f(uint(x))
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", impl.name, x, got, expected)
			}
		}
		gen := random.NewUintGenerator()
		for i := 0; i < 512; i++ {
			x := gen.Next()
			expected := OnesCountNaive(x)
			got := impl.f(x)
			if got != expected {
				t.Errorf("%s(%d) = %d, expected %d", impl.name, x, got, expected)
			}
		}
	}
}
