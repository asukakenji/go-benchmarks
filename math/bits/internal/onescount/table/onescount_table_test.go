package table_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

func TestOnesCountAllNaive(t *testing.T) {
	onescount.NaiveTest(t, table.OnesCount)
	onescount.NaiveTest8(t, table.OnesCount8)
	onescount.NaiveTest16(t, table.OnesCount16)
	onescount.NaiveTest32(t, table.OnesCount32)
	onescount.NaiveTest64(t, table.OnesCount64)
}

func TestOnesCount8(t *testing.T) {
	cases := []struct {
		x        uint8
		expected int
	}{
		{0x00, 0}, {0x11, 2}, {0x22, 2}, {0x33, 4},
		{0x44, 2}, {0x55, 4}, {0x66, 4}, {0x77, 6},
		{0x88, 2}, {0x99, 4}, {0xaa, 4}, {0xbb, 6},
		{0xcc, 4}, {0xdd, 6}, {0xee, 6}, {0xff, 8},

		{0x01, 1}, {0x23, 3}, {0x45, 3}, {0x67, 5},
		{0x89, 3}, {0xab, 5}, {0xcd, 5}, {0xef, 7},
		{0x10, 1}, {0x32, 3}, {0x54, 3}, {0x76, 5},
		{0x98, 3}, {0xba, 5}, {0xdc, 5}, {0xfe, 7},
	}
	for _, c := range cases {
		got := table.OnesCount8(c.x)
		if got != c.expected {
			t.Errorf("table.OnesCount8(%d) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestOnesCount16(t *testing.T) {
	cases := []struct {
		x        uint16
		expected int
	}{
		{0x0000, 0}, {0x1111, 4}, {0x2222, 4}, {0x3333, 8},
		{0x4444, 4}, {0x5555, 8}, {0x6666, 8}, {0x7777, 12},
		{0x8888, 4}, {0x9999, 8}, {0xaaaa, 8}, {0xbbbb, 12},
		{0xcccc, 8}, {0xdddd, 12}, {0xeeee, 12}, {0xffff, 16},

		{0x0123, 4}, {0x4567, 8}, {0x89ab, 8}, {0xcdef, 12},
		{0x3210, 4}, {0x7654, 8}, {0xba98, 8}, {0xfedc, 12},
	}
	for _, c := range cases {
		got := table.OnesCount16(c.x)
		if got != c.expected {
			t.Errorf("table.OnesCount16(%d) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestOnesCount32(t *testing.T) {
	cases := []struct {
		x        uint32
		expected int
	}{
		{0x00000000, 0}, {0x11111111, 8}, {0x22222222, 8}, {0x33333333, 16},
		{0x44444444, 8}, {0x55555555, 16}, {0x66666666, 16}, {0x77777777, 24},
		{0x88888888, 8}, {0x99999999, 16}, {0xaaaaaaaa, 16}, {0xbbbbbbbb, 24},
		{0xcccccccc, 16}, {0xdddddddd, 24}, {0xeeeeeeee, 24}, {0xffffffff, 32},

		{0x01234567, 12}, {0x89abcdef, 20},
		{0x76543210, 12}, {0xfedcba98, 20},
	}
	for _, c := range cases {
		got := table.OnesCount32(c.x)
		if got != c.expected {
			t.Errorf("table.OnesCount32(%d) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestOnesCount64(t *testing.T) {
	cases := []struct {
		x        uint64
		expected int
	}{
		{0x0000000000000000, 0}, {0x1111111111111111, 16},
		{0x2222222222222222, 16}, {0x3333333333333333, 32},
		{0x4444444444444444, 16}, {0x5555555555555555, 32},
		{0x6666666666666666, 32}, {0x7777777777777777, 48},
		{0x8888888888888888, 16}, {0x9999999999999999, 32},
		{0xaaaaaaaaaaaaaaaa, 32}, {0xbbbbbbbbbbbbbbbb, 48},
		{0xcccccccccccccccc, 32}, {0xdddddddddddddddd, 48},
		{0xeeeeeeeeeeeeeeee, 48}, {0xffffffffffffffff, 64},

		{0x0123456789abcdef, 32},
		{0xfedcba9876543210, 32},
	}
	for _, c := range cases {
		got := table.OnesCount64(c.x)
		if got != c.expected {
			t.Errorf("table.OnesCount64(%d) = %d, expected %d", c.x, got, c.expected)
		}
	}
}
