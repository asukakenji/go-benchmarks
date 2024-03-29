package tcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks"
)

func TestLeadingZeros(t *testing.T, leadingZeros func(uint) int) {
	cases := []struct {
		x        uint
		expected int
	}{
		{0, benchmarks.SizeOfUintInBits},
		{1 << 0, benchmarks.SizeOfUintInBits - 1},
		{1 << 1, benchmarks.SizeOfUintInBits - 2},
		{1 << 2, benchmarks.SizeOfUintInBits - 3},
		{1 << 3, benchmarks.SizeOfUintInBits - 4},
		{1 << 4, benchmarks.SizeOfUintInBits - 5},
		{1 << 5, benchmarks.SizeOfUintInBits - 6},
		{1 << 6, benchmarks.SizeOfUintInBits - 7},
		{1 << 7, benchmarks.SizeOfUintInBits - 8},
		{1 << 8, benchmarks.SizeOfUintInBits - 9},
		{1 << 9, benchmarks.SizeOfUintInBits - 10},
		{1 << 10, benchmarks.SizeOfUintInBits - 11},
		{1 << 11, benchmarks.SizeOfUintInBits - 12},
		{1 << 12, benchmarks.SizeOfUintInBits - 13},
		{1 << 13, benchmarks.SizeOfUintInBits - 14},
		{1 << 14, benchmarks.SizeOfUintInBits - 15},
		{1 << 15, benchmarks.SizeOfUintInBits - 16},
		{1 << 16, benchmarks.SizeOfUintInBits - 17},
		{1 << 17, benchmarks.SizeOfUintInBits - 18},
		{1 << 18, benchmarks.SizeOfUintInBits - 19},
		{1 << 19, benchmarks.SizeOfUintInBits - 20},
		{1 << 20, benchmarks.SizeOfUintInBits - 21},
		{1 << 21, benchmarks.SizeOfUintInBits - 22},
		{1 << 22, benchmarks.SizeOfUintInBits - 23},
		{1 << 23, benchmarks.SizeOfUintInBits - 24},
		{1 << 24, benchmarks.SizeOfUintInBits - 25},
		{1 << 25, benchmarks.SizeOfUintInBits - 26},
		{1 << 26, benchmarks.SizeOfUintInBits - 27},
		{1 << 27, benchmarks.SizeOfUintInBits - 28},
		{1 << 28, benchmarks.SizeOfUintInBits - 29},
		{1 << 29, benchmarks.SizeOfUintInBits - 30},
		{1 << 30, benchmarks.SizeOfUintInBits - 31},
		{1 << 31, benchmarks.SizeOfUintInBits - 32},
	}
	for _, c := range cases {
		got := leadingZeros(c.x)
		if got != c.expected {
			t.Errorf("LeadingZeros(0x%x) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestLeadingZeros8(t *testing.T, leadingZeros8 func(uint8) int) {
	cases := []struct {
		x        uint8
		expected int
	}{
		{0x00, 8}, {0x11, 3}, {0x22, 2}, {0x33, 2},
		{0x44, 1}, {0x55, 1}, {0x66, 1}, {0x77, 1},
		{0x88, 0}, {0x99, 0}, {0xaa, 0}, {0xbb, 0},
		{0xcc, 0}, {0xdd, 0}, {0xee, 0}, {0xff, 0},

		{0x01, 7}, {0x23, 2}, {0x45, 1}, {0x67, 1},
		{0x89, 0}, {0xab, 0}, {0xcd, 0}, {0xef, 0},
		{0x10, 3}, {0x32, 2}, {0x54, 1}, {0x76, 1},
		{0x98, 0}, {0xba, 0}, {0xdc, 0}, {0xfe, 0},

		{1 << 0, 7}, {1 << 1, 6}, {1 << 2, 5}, {1 << 3, 4},
		{1 << 4, 3}, {1 << 5, 2}, {1 << 6, 1}, {1 << 7, 0},
	}
	for _, c := range cases {
		got := leadingZeros8(c.x)
		if got != c.expected {
			t.Errorf("LeadingZeros8(0x%02x) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestLeadingZeros16(t *testing.T, leadingZeros16 func(uint16) int) {
	cases := []struct {
		x        uint16
		expected int
	}{
		{0x0000, 16}, {0x1111, 3}, {0x2222, 2}, {0x3333, 2},
		{0x4444, 1}, {0x5555, 1}, {0x6666, 1}, {0x7777, 1},
		{0x8888, 0}, {0x9999, 0}, {0xaaaa, 0}, {0xbbbb, 0},
		{0xcccc, 0}, {0xdddd, 0}, {0xeeee, 0}, {0xffff, 0},

		{0x0123, 7}, {0x4567, 1}, {0x89ab, 0}, {0xcdef, 0},
		{0x3210, 2}, {0x7654, 1}, {0xba98, 0}, {0xfedc, 0},

		{1 << 0, 15}, {1 << 1, 14}, {1 << 2, 13}, {1 << 3, 12},
		{1 << 4, 11}, {1 << 5, 10}, {1 << 6, 9}, {1 << 7, 8},
		{1 << 8, 7}, {1 << 9, 6}, {1 << 10, 5}, {1 << 11, 4},
		{1 << 12, 3}, {1 << 13, 2}, {1 << 14, 1}, {1 << 15, 0},
	}
	for _, c := range cases {
		got := leadingZeros16(c.x)
		if got != c.expected {
			t.Errorf("LeadingZeros16(0x%04x) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestLeadingZeros32(t *testing.T, leadingZeros32 func(uint32) int) {
	cases := []struct {
		x        uint32
		expected int
	}{
		{0x00000000, 32}, {0x11111111, 3}, {0x22222222, 2}, {0x33333333, 2},
		{0x44444444, 1}, {0x55555555, 1}, {0x66666666, 1}, {0x77777777, 1},
		{0x88888888, 0}, {0x99999999, 0}, {0xaaaaaaaa, 0}, {0xbbbbbbbb, 0},
		{0xcccccccc, 0}, {0xdddddddd, 0}, {0xeeeeeeee, 0}, {0xffffffff, 0},

		{0x01234567, 7}, {0x89abcdef, 0},
		{0x76543210, 1}, {0xfedcba98, 0},

		{1 << 0, 31}, {1 << 1, 30}, {1 << 2, 29}, {1 << 3, 28},
		{1 << 4, 27}, {1 << 5, 26}, {1 << 6, 25}, {1 << 7, 24},
		{1 << 8, 23}, {1 << 9, 22}, {1 << 10, 21}, {1 << 11, 20},
		{1 << 12, 19}, {1 << 13, 18}, {1 << 14, 17}, {1 << 15, 16},

		{1 << 16, 15}, {1 << 17, 14}, {1 << 18, 13}, {1 << 19, 12},
		{1 << 20, 11}, {1 << 21, 10}, {1 << 22, 9}, {1 << 23, 8},
		{1 << 24, 7}, {1 << 25, 6}, {1 << 26, 5}, {1 << 27, 4},
		{1 << 28, 3}, {1 << 29, 2}, {1 << 30, 1}, {1 << 31, 0},
	}
	for _, c := range cases {
		got := leadingZeros32(c.x)
		if got != c.expected {
			t.Errorf("LeadingZeros32(0x%08x) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestLeadingZeros64(t *testing.T, leadingZeros64 func(uint64) int) {
	cases := []struct {
		x        uint64
		expected int
	}{
		{0x0000000000000000, 64}, {0x1111111111111111, 3},
		{0x2222222222222222, 2}, {0x3333333333333333, 2},
		{0x4444444444444444, 1}, {0x5555555555555555, 1},
		{0x6666666666666666, 1}, {0x7777777777777777, 1},
		{0x8888888888888888, 0}, {0x9999999999999999, 0},
		{0xaaaaaaaaaaaaaaaa, 0}, {0xbbbbbbbbbbbbbbbb, 0},
		{0xcccccccccccccccc, 0}, {0xdddddddddddddddd, 0},
		{0xeeeeeeeeeeeeeeee, 0}, {0xffffffffffffffff, 0},

		{0x0123456789abcdef, 7},
		{0xfedcba9876543210, 0},

		{1 << 0, 63}, {1 << 1, 62}, {1 << 2, 61}, {1 << 3, 60},
		{1 << 4, 59}, {1 << 5, 58}, {1 << 6, 57}, {1 << 7, 56},
		{1 << 8, 55}, {1 << 9, 54}, {1 << 10, 53}, {1 << 11, 52},
		{1 << 12, 51}, {1 << 13, 50}, {1 << 14, 49}, {1 << 15, 48},

		{1 << 16, 47}, {1 << 17, 46}, {1 << 18, 45}, {1 << 19, 44},
		{1 << 20, 43}, {1 << 21, 42}, {1 << 22, 41}, {1 << 23, 40},
		{1 << 24, 39}, {1 << 25, 38}, {1 << 26, 37}, {1 << 27, 36},
		{1 << 28, 35}, {1 << 29, 34}, {1 << 30, 33}, {1 << 31, 32},

		{1 << 32, 31}, {1 << 33, 30}, {1 << 34, 29}, {1 << 35, 28},
		{1 << 36, 27}, {1 << 37, 26}, {1 << 38, 25}, {1 << 39, 24},
		{1 << 40, 23}, {1 << 41, 22}, {1 << 42, 21}, {1 << 43, 20},
		{1 << 44, 19}, {1 << 45, 18}, {1 << 46, 17}, {1 << 47, 16},

		{1 << 48, 15}, {1 << 49, 14}, {1 << 50, 13}, {1 << 51, 12},
		{1 << 52, 11}, {1 << 53, 10}, {1 << 54, 9}, {1 << 55, 8},
		{1 << 56, 7}, {1 << 57, 6}, {1 << 58, 5}, {1 << 59, 4},
		{1 << 60, 3}, {1 << 61, 2}, {1 << 62, 1}, {1 << 63, 0},
	}
	for _, c := range cases {
		got := leadingZeros64(c.x)
		if got != c.expected {
			t.Errorf("LeadingZeros64(0x%016x) = %d, expected %d", c.x, got, c.expected)
		}
	}
}

func TestLeadingZerosPtr(t *testing.T, leadingZerosPtr func(uintptr) int) {
	cases := []struct {
		x        uintptr
		expected int
	}{
		{0, benchmarks.SizeOfUintInBits},
		{1 << 0, benchmarks.SizeOfUintInBits - 1},
		{1 << 1, benchmarks.SizeOfUintInBits - 2},
		{1 << 2, benchmarks.SizeOfUintInBits - 3},
		{1 << 3, benchmarks.SizeOfUintInBits - 4},
		{1 << 4, benchmarks.SizeOfUintInBits - 5},
		{1 << 5, benchmarks.SizeOfUintInBits - 6},
		{1 << 6, benchmarks.SizeOfUintInBits - 7},
		{1 << 7, benchmarks.SizeOfUintInBits - 8},
		{1 << 8, benchmarks.SizeOfUintInBits - 9},
		{1 << 9, benchmarks.SizeOfUintInBits - 10},
		{1 << 10, benchmarks.SizeOfUintInBits - 11},
		{1 << 11, benchmarks.SizeOfUintInBits - 12},
		{1 << 12, benchmarks.SizeOfUintInBits - 13},
		{1 << 13, benchmarks.SizeOfUintInBits - 14},
		{1 << 14, benchmarks.SizeOfUintInBits - 15},
		{1 << 15, benchmarks.SizeOfUintInBits - 16},
		{1 << 16, benchmarks.SizeOfUintInBits - 17},
		{1 << 17, benchmarks.SizeOfUintInBits - 18},
		{1 << 18, benchmarks.SizeOfUintInBits - 19},
		{1 << 19, benchmarks.SizeOfUintInBits - 20},
		{1 << 20, benchmarks.SizeOfUintInBits - 21},
		{1 << 21, benchmarks.SizeOfUintInBits - 22},
		{1 << 22, benchmarks.SizeOfUintInBits - 23},
		{1 << 23, benchmarks.SizeOfUintInBits - 24},
		{1 << 24, benchmarks.SizeOfUintInBits - 25},
		{1 << 25, benchmarks.SizeOfUintInBits - 26},
		{1 << 26, benchmarks.SizeOfUintInBits - 27},
		{1 << 27, benchmarks.SizeOfUintInBits - 28},
		{1 << 28, benchmarks.SizeOfUintInBits - 29},
		{1 << 29, benchmarks.SizeOfUintInBits - 30},
		{1 << 30, benchmarks.SizeOfUintInBits - 31},
		{1 << 31, benchmarks.SizeOfUintInBits - 32},
	}
	for _, c := range cases {
		got := leadingZerosPtr(c.x)
		if got != c.expected {
			t.Errorf("LeadingZerosPtr(0x%x) = %d, expected %d", c.x, got, c.expected)
		}
	}
}
