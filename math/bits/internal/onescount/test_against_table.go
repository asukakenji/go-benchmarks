package onescount

import (
	"math/rand"
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/table"
)

var pop8tab = [256]uint8{
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

type OnesCountFuncs struct {
	Uint   func(uint) int
	Uint8  func(uint8) int
	Uint16 func(uint16) int
	Uint32 func(uint32) int
	Uint64 func(uint64) int
}

func BasicTableTest(t *testing.T, onesCountFuncs *OnesCountFuncs) {
	if onesCountFuncs.Uint != nil {
		for x, expected := range pop8tab {
			got := onesCountFuncs.Uint(uint(x))
			if got != int(expected) {
				t.Errorf("OnesCount(0x%x) = %d, expected %d", x, got, expected)
			}
		}
	}
	if onesCountFuncs.Uint8 != nil {
		for x, expected := range pop8tab {
			got := onesCountFuncs.Uint8(uint8(x))
			if got != int(expected) {
				t.Errorf("OnesCount8(0x%02x) = %d, expected %d", x, got, expected)
			}
		}
	}
	if onesCountFuncs.Uint16 != nil {
		for x, expected := range pop8tab {
			got := onesCountFuncs.Uint16(uint16(x))
			if got != int(expected) {
				t.Errorf("OnesCount16(0x%04x) = %d, expected %d", x, got, expected)
			}
		}
	}
	if onesCountFuncs.Uint32 != nil {
		for x, expected := range pop8tab {
			got := onesCountFuncs.Uint32(uint32(x))
			if got != int(expected) {
				t.Errorf("OnesCount32(0x%08x) = %d, expected %d", x, got, expected)
			}
		}
	}
	if onesCountFuncs.Uint64 != nil {
		for x, expected := range pop8tab {
			got := onesCountFuncs.Uint64(uint64(x))
			if got != int(expected) {
				t.Errorf("OnesCount64(0x%016x) = %d, expected %d", x, got, expected)
			}
		}
		cases := []struct {
			x        uint64
			expected int
		}{
			{0xfffffffffffffff7, 63},
			{0xfffffffffffffffb, 63},
			{0xfffffffffffffffd, 63},
			{0xfffffffffffffffe, 63},
			{0xffffffffffffff7f, 63},
			{0xffffffffffffffbf, 63},
			{0xffffffffffffffdf, 63},
			{0xffffffffffffffef, 63},
			{0xfffffffffffff7ff, 63},
			{0xfffffffffffffbff, 63},
			{0xfffffffffffffdff, 63},
			{0xfffffffffffffeff, 63},
			{0xffffffffffff7fff, 63},
			{0xffffffffffffbfff, 63},
			{0xffffffffffffdfff, 63},
			{0xffffffffffffefff, 63},
			{0xfffffffffff7ffff, 63},
			{0xfffffffffffbffff, 63},
			{0xfffffffffffdffff, 63},
			{0xfffffffffffeffff, 63},
			{0xffffffffff7fffff, 63},
			{0xffffffffffbfffff, 63},
			{0xffffffffffdfffff, 63},
			{0xffffffffffefffff, 63},
			{0xfffffffff7ffffff, 63},
			{0xfffffffffbffffff, 63},
			{0xfffffffffdffffff, 63},
			{0xfffffffffeffffff, 63},
			{0xffffffff7fffffff, 63},
			{0xffffffffbfffffff, 63},
			{0xffffffffdfffffff, 63},
			{0xffffffffefffffff, 63},
			{0xfffffff7ffffffff, 63},
			{0xfffffffbffffffff, 63},
			{0xfffffffdffffffff, 63},
			{0xfffffffeffffffff, 63},
			{0xffffff7fffffffff, 63},
			{0xffffffbfffffffff, 63},
			{0xffffffdfffffffff, 63},
			{0xffffffefffffffff, 63},
			{0xfffff7ffffffffff, 63},
			{0xfffffbffffffffff, 63},
			{0xfffffdffffffffff, 63},
			{0xfffffeffffffffff, 63},
			{0xffff7fffffffffff, 63},
			{0xffffbfffffffffff, 63},
			{0xffffdfffffffffff, 63},
			{0xffffefffffffffff, 63},
			{0xfff7ffffffffffff, 63},
			{0xfffbffffffffffff, 63},
			{0xfffdffffffffffff, 63},
			{0xfffeffffffffffff, 63},
			{0xff7fffffffffffff, 63},
			{0xffbfffffffffffff, 63},
			{0xffdfffffffffffff, 63},
			{0xffefffffffffffff, 63},
			{0xf7ffffffffffffff, 63},
			{0xfbffffffffffffff, 63},
			{0xfdffffffffffffff, 63},
			{0xfeffffffffffffff, 63},
			{0x7fffffffffffffff, 63},
			{0xbfffffffffffffff, 63},
			{0xdfffffffffffffff, 63},
			{0xefffffffffffffff, 63},
			{0xffffffffffffffff, 64},
		}
		for _, c := range cases {
			got := onesCountFuncs.Uint64(c.x)
			if got != c.expected {
				t.Errorf("OnesCount64(0x%016x) = %d, expected %d", c.x, got, c.expected)
			}
		}
	}
}

func TableTest(t *testing.T, onesCountFunc func(uint) int) {
	for i := 0; i < 256; i++ {
		x := uint(rand.Uint64())
		expected := table.OnesCountConcept(x)
		got := onesCountFunc(x)
		if got != expected {
			t.Errorf("OnesCount(0x%x) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest7(t *testing.T, onesCount7Func func(uint8) int) {
	for i := 0; i < 256; i++ {
		var x uint8
		for {
			x = uint8(rand.Uint64())
			if x <= 0x7f {
				break
			}
		}
		expected := table.OnesCount8(x)
		got := onesCount7Func(x)
		if got != expected {
			t.Errorf("OnesCount7(0x%02x) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest8(t *testing.T, onesCount8Func func(uint8) int) {
	for i := 0; i < 256; i++ {
		x := uint8(rand.Uint64())
		expected := table.OnesCount8(x)
		got := onesCount8Func(x)
		if got != expected {
			t.Errorf("OnesCount8(0x%02x) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest15(t *testing.T, onesCount15Func func(uint16) int) {
	for i := 0; i < 256; i++ {
		var x uint16
		for {
			x = uint16(rand.Uint64())
			if x <= 0x7fff {
				break
			}
		}
		expected := table.OnesCount16(x)
		got := onesCount15Func(x)
		if got != expected {
			t.Errorf("OnesCount15(0x%04x) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest16(t *testing.T, onesCount16Func func(uint16) int) {
	for i := 0; i < 256; i++ {
		x := uint16(rand.Uint64())
		expected := table.OnesCount16(x)
		got := onesCount16Func(x)
		if got != expected {
			t.Errorf("OnesCount16(0x%04x) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest32(t *testing.T, onesCount32Func func(uint32) int) {
	for i := 0; i < 256; i++ {
		x := uint32(rand.Uint64())
		expected := table.OnesCount32(x)
		got := onesCount32Func(x)
		if got != expected {
			t.Errorf("OnesCount32(0x%08x) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest64(t *testing.T, onesCount64Func func(uint64) int) {
	for i := 0; i < 256; i++ {
		x := rand.Uint64()
		expected := table.OnesCount64(x)
		got := onesCount64Func(x)
		if got != expected {
			t.Errorf("OnesCount64(0x%016x) = %d, expected %d", x, got, expected)
		}
	}
}
