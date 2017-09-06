package leadingzeros

import (
	"math/rand"
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/naive"
)

func NaiveTest(t *testing.T, leadingZerosFunc func(uint) int) {
	for i := 0; i < 256; i++ {
		x := uint(rand.Uint64())
		expected := naive.LeadingZeros(x)
		got := leadingZerosFunc(x)
		if got != expected {
			t.Errorf("LeadingZeros(0x%x) = %d, expected %d", x, got, expected)
		}
	}
}

func NaiveTest8(t *testing.T, leadingZeros8Func func(uint8) int) {
	for i := 0; i < 256; i++ {
		x := uint8(rand.Uint64())
		expected := naive.LeadingZeros8(x)
		got := leadingZeros8Func(x)
		if got != expected {
			t.Errorf("LeadingZeros8(0x%02x) = %d, expected %d", x, got, expected)
		}
	}
}

func NaiveTest16(t *testing.T, leadingZeros16Func func(uint16) int) {
	for i := 0; i < 256; i++ {
		x := uint16(rand.Uint64())
		expected := naive.LeadingZeros16(x)
		got := leadingZeros16Func(x)
		if got != expected {
			t.Errorf("LeadingZeros16(0x%04x) = %d, expected %d", x, got, expected)
		}
	}
}

func NaiveTest32(t *testing.T, leadingZeros32Func func(uint32) int) {
	for i := 0; i < 256; i++ {
		x := uint32(rand.Uint64())
		expected := naive.LeadingZeros32(x)
		got := leadingZeros32Func(x)
		if got != expected {
			t.Errorf("LeadingZeros32(0x%08x) = %d, expected %d", x, got, expected)
		}
	}
}

func NaiveTest64(t *testing.T, leadingZeros64Func func(uint64) int) {
	for i := 0; i < 256; i++ {
		x := rand.Uint64()
		expected := naive.LeadingZeros64(x)
		got := leadingZeros64Func(x)
		if got != expected {
			t.Errorf("LeadingZeros64(0x%016x) = %d, expected %d", x, got, expected)
		}
	}
}
