package tcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks"
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

func TestBasic(t *testing.T, leadingZeros func(uint) int) {
	for x, v := range nlz8tab {
		expected := int(v) + benchmarks.SizeOfUintInBits - 8
		got := leadingZeros(uint(x))
		if got != expected {
			t.Errorf("LeadingZeros(0x%x) = %d, expected %d", x, got, expected)
		}
	}
}

func TestBasic8(t *testing.T, leadingZeros func(uint8) int) {
	for x, v := range nlz8tab {
		expected := int(v)
		got := leadingZeros(uint8(x))
		if got != expected {
			t.Errorf("LeadingZeros8(0x%02x) = %d, expected %d", x, got, expected)
		}
	}
}

func TestBasic16(t *testing.T, leadingZeros func(uint16) int) {
	for x, v := range nlz8tab {
		expected := int(v) + 8
		got := leadingZeros(uint16(x))
		if got != expected {
			t.Errorf("LeadingZeros16(0x%04x) = %d, expected %d", x, got, expected)
		}
	}
}

func TestBasic32(t *testing.T, leadingZeros func(uint32) int) {
	for x, v := range nlz8tab {
		expected := int(v) + 24
		got := leadingZeros(uint32(x))
		if got != expected {
			t.Errorf("LeadingZeros32(0x%08x) = %d, expected %d", x, got, expected)
		}
	}
}

func TestBasic64(t *testing.T, leadingZeros func(uint64) int) {
	for x, v := range nlz8tab {
		expected := int(v) + 56
		got := leadingZeros(uint64(x))
		if got != expected {
			t.Errorf("LeadingZeros64(0x%016x) = %d, expected %d", x, got, expected)
		}
	}
}

func TestBasicPtr(t *testing.T, leadingZeros func(uintptr) int) {
	for x, v := range nlz8tab {
		expected := int(v) + benchmarks.SizeOf[uintptr]() - 8
		got := leadingZeros(uintptr(x))
		if got != expected {
			t.Errorf("LeadingZeros(0x%x) = %d, expected %d", x, got, expected)
		}
	}
}
