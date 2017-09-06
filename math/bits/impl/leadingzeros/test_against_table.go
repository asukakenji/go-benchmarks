package leadingzeros

import (
	"math/rand"
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/table"
)

const (
	uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
)

var nlz8tab = [256]uint8{
	8, 7, 6, 6, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,

	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,

	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

type LeadingZerosFuncs struct {
	Uint   func(uint) int
	Uint8  func(uint8) int
	Uint16 func(uint16) int
	Uint32 func(uint32) int
	Uint64 func(uint64) int
}

func BasicTableTest(t *testing.T, leadingZerosFuncs *LeadingZerosFuncs) {
	if leadingZerosFuncs.Uint != nil {
		for x, v := range nlz8tab {
			expected := int(v + uintSize - 8)
			got := leadingZerosFuncs.Uint(uint(x))
			if got != expected {
				t.Errorf("LeadingZeros(0x%x) = %d, expected %d", x, got, expected)
			}
		}
	}
	if leadingZerosFuncs.Uint8 != nil {
		for x, v := range nlz8tab {
			expected := int(v)
			got := leadingZerosFuncs.Uint8(uint8(x))
			if got != expected {
				t.Errorf("LeadingZeros8(0x%02x) = %d, expected %d", x, got, expected)
			}
		}
	}
	if leadingZerosFuncs.Uint16 != nil {
		for x, v := range nlz8tab {
			expected := int(v + 8)
			got := leadingZerosFuncs.Uint16(uint16(x))
			if got != expected {
				t.Errorf("LeadingZeros16(0x%04x) = %d, expected %d", x, got, expected)
			}
		}
	}
	if leadingZerosFuncs.Uint32 != nil {
		for x, v := range nlz8tab {
			expected := int(v + 24)
			got := leadingZerosFuncs.Uint32(uint32(x))
			if got != expected {
				t.Errorf("LeadingZeros32(0x%08x) = %d, expected %d", x, got, expected)
			}
		}
	}
	if leadingZerosFuncs.Uint64 != nil {
		for x, v := range nlz8tab {
			expected := int(v + 56)
			got := leadingZerosFuncs.Uint64(uint64(x))
			if got != expected {
				t.Errorf("LeadingZeros64(0x%016x) = %d, expected %d", x, got, expected)
			}
		}
	}
}

func TableTest(t *testing.T, leadingZerosFunc func(uint) int) {
	for i := 0; i < 256; i++ {
		x := uint(rand.Uint64())
		expected := table.LeadingZerosConcept(x)
		got := leadingZerosFunc(x)
		if got != expected {
			t.Errorf("LeadingZeros(0x%x) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest8(t *testing.T, leadingZeros8Func func(uint8) int) {
	for i := 0; i < 256; i++ {
		x := uint8(rand.Uint64())
		expected := table.LeadingZeros8(x)
		got := leadingZeros8Func(x)
		if got != expected {
			t.Errorf("LeadingZeros8(0x%02x) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest16(t *testing.T, leadingZeros16Func func(uint16) int) {
	for i := 0; i < 256; i++ {
		x := uint16(rand.Uint64())
		expected := table.LeadingZeros16(x)
		got := leadingZeros16Func(x)
		if got != expected {
			t.Errorf("LeadingZeros16(0x%04x) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest32(t *testing.T, leadingZeros32Func func(uint32) int) {
	for i := 0; i < 256; i++ {
		x := uint32(rand.Uint64())
		expected := table.LeadingZeros32(x)
		got := leadingZeros32Func(x)
		if got != expected {
			t.Errorf("LeadingZeros32(0x%08x) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest64(t *testing.T, leadingZeros64Func func(uint64) int) {
	for i := 0; i < 256; i++ {
		x := rand.Uint64()
		expected := table.LeadingZeros64(x)
		got := leadingZeros64Func(x)
		if got != expected {
			t.Errorf("LeadingZeros64(0x%016x) = %d, expected %d", x, got, expected)
		}
	}
}
