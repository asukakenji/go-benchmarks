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

func testAgainst8(t *testing.T, referenceLeadingZeros, targetLeadingZeros func(uint8) int) {
	for i := 0; i < 256; i++ {
		x := uint8(rand.Uint64())
		expected := referenceLeadingZeros(x)
		got := targetLeadingZeros(x)
		if got != expected {
			t.Errorf("LeadingZeros8(0x%02x) = %d, expected %d", x, got, expected)
		}
	}
}

func testAgainst16(t *testing.T, referenceLeadingZeros, targetLeadingZeros func(uint16) int) {
	for i := 0; i < 256; i++ {
		x := uint16(rand.Uint64())
		expected := referenceLeadingZeros(x)
		got := targetLeadingZeros(x)
		if got != expected {
			t.Errorf("LeadingZeros16(0x%04x) = %d, expected %d", x, got, expected)
		}
	}
}

func testAgainst32(t *testing.T, referenceLeadingZeros, targetLeadingZeros func(uint32) int) {
	for i := 0; i < 256; i++ {
		x := uint32(rand.Uint64())
		expected := referenceLeadingZeros(x)
		got := targetLeadingZeros(x)
		if got != expected {
			t.Errorf("LeadingZeros32(0x%08x) = %d, expected %d", x, got, expected)
		}
	}
}

func testAgainst64(t *testing.T, referenceLeadingZeros, targetLeadingZeros func(uint64) int) {
	for i := 0; i < 256; i++ {
		x := rand.Uint64()
		expected := referenceLeadingZeros(x)
		got := targetLeadingZeros(x)
		if got != expected {
			t.Errorf("LeadingZeros64(0x%016x) = %d, expected %d", x, got, expected)
		}
	}
}

func testAgainstPtr(t *testing.T, referenceLeadingZeros, targetLeadingZeros func(uintptr) int) {
	for i := 0; i < 256; i++ {
		x := uintptr(rand.Uint64())
		expected := referenceLeadingZeros(x)
		got := targetLeadingZeros(x)
		if got != expected {
			t.Errorf("LeadingZerosPtr(0x%x) = %d, expected %d", x, got, expected)
		}
	}
}
