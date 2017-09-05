package onescount

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
)

func TestOnesCount32All(t *testing.T) {
	implementations := []struct {
		name string
		f    func(uint32) int
	}{
		{"OnesCount32Pop2", OnesCount32Pop2},
		{"OnesCount32Pop2Alt", OnesCount32Pop2Alt},
		{"OnesCount32Pop3", OnesCount32Pop3},
		{"OnesCount32Pop5", OnesCount32Pop5},
		{"OnesCount32Hakmem", OnesCount32Hakmem},
		{"OnesCount32HakmemUnrolled", OnesCount32HakmemUnrolled},
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
		for _, x := range cases {
			expected := naive.OnesCount32(uint32(x))
			got := impl.f(uint32(x))
			if got != expected {
				t.Errorf("%s(0x%08x) = %d, expected %d", impl.name, x, got, expected)
			}
		}
		gen := randomsupplier.NewUint32()
		for i := 0; i < 512; i++ {
			x := gen.Next()
			expected := naive.OnesCount32(x)
			got := impl.f(x)
			if got != expected {
				t.Errorf("%s(0x%08x) = %d, expected %d", impl.name, x, got, expected)
			}
		}
	}
}
