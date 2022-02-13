package tcommon

import (
	"math/rand"
	"testing"
)

func testAgainst(t *testing.T, referenceLeadingZeros, targetLeadingZeros func(uint) int) {
	for i := 0; i < 256; i++ {
		x := uint(rand.Uint64())
		expected := referenceLeadingZeros(x)
		got := targetLeadingZeros(x)
		if got != expected {
			t.Errorf("LeadingZeros(0x%x) = %d, expected %d", x, got, expected)
		}
	}
}

func testAgainst8(t *testing.T, referenceLeadingZeros8, targetLeadingZeros8 func(uint8) int) {
	for i := 0; i < 256; i++ {
		x := uint8(rand.Uint64())
		expected := referenceLeadingZeros8(x)
		got := targetLeadingZeros8(x)
		if got != expected {
			t.Errorf("LeadingZeros8(0x%02x) = %d, expected %d", x, got, expected)
		}
	}
}

func testAgainst16(t *testing.T, referenceLeadingZeros16, targetLeadingZeros16 func(uint16) int) {
	for i := 0; i < 256; i++ {
		x := uint16(rand.Uint64())
		expected := referenceLeadingZeros16(x)
		got := targetLeadingZeros16(x)
		if got != expected {
			t.Errorf("LeadingZeros16(0x%04x) = %d, expected %d", x, got, expected)
		}
	}
}

func testAgainst32(t *testing.T, referenceLeadingZeros32, targetLeadingZeros32 func(uint32) int) {
	for i := 0; i < 256; i++ {
		x := uint32(rand.Uint64())
		expected := referenceLeadingZeros32(x)
		got := targetLeadingZeros32(x)
		if got != expected {
			t.Errorf("LeadingZeros32(0x%08x) = %d, expected %d", x, got, expected)
		}
	}
}

func testAgainst64(t *testing.T, referenceLeadingZeros64, targetLeadingZeros64 func(uint64) int) {
	for i := 0; i < 256; i++ {
		x := rand.Uint64()
		expected := referenceLeadingZeros64(x)
		got := targetLeadingZeros64(x)
		if got != expected {
			t.Errorf("LeadingZeros64(0x%016x) = %d, expected %d", x, got, expected)
		}
	}
}

func testAgainstPtr(t *testing.T, referenceLeadingZerosPtr, targetLeadingZerosPtr func(uintptr) int) {
	for i := 0; i < 256; i++ {
		x := uintptr(rand.Uint64())
		expected := referenceLeadingZerosPtr(x)
		got := targetLeadingZerosPtr(x)
		if got != expected {
			t.Errorf("LeadingZerosPtr(0x%x) = %d, expected %d", x, got, expected)
		}
	}
}
