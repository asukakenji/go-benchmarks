package onescount

import (
	"math/rand"
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/naive"
)

func NaiveTest(t *testing.T, onesCountFunc func(uint) int) {
	for i := 0; i < 256; i++ {
		x := uint(rand.Uint64())
		expected := naive.OnesCount(x)
		got := onesCountFunc(x)
		if got != expected {
			t.Errorf("OnesCount(0x%x) = %d, expected %d", x, got, expected)
		}
	}
}

func NaiveTest8(t *testing.T, onesCount8Func func(uint8) int) {
	for i := 0; i < 256; i++ {
		x := uint8(rand.Uint64())
		expected := naive.OnesCount8(x)
		got := onesCount8Func(x)
		if got != expected {
			t.Errorf("OnesCount8(0x%02x) = %d, expected %d", x, got, expected)
		}
	}
}

func NaiveTest16(t *testing.T, onesCount16Func func(uint16) int) {
	for i := 0; i < 256; i++ {
		x := uint16(rand.Uint64())
		expected := naive.OnesCount16(x)
		got := onesCount16Func(x)
		if got != expected {
			t.Errorf("OnesCount16(0x%04x) = %d, expected %d", x, got, expected)
		}
	}
}

func NaiveTest32(t *testing.T, onesCount32Func func(uint32) int) {
	for i := 0; i < 256; i++ {
		x := uint32(rand.Uint64())
		expected := naive.OnesCount32(x)
		got := onesCount32Func(x)
		if got != expected {
			t.Errorf("OnesCount32(0x%08x) = %d, expected %d", x, got, expected)
		}
	}
}

func NaiveTest64(t *testing.T, onesCount64Func func(uint64) int) {
	for i := 0; i < 256; i++ {
		x := rand.Uint64()
		expected := naive.OnesCount64(x)
		got := onesCount64Func(x)
		if got != expected {
			t.Errorf("OnesCount64(0x%016x) = %d, expected %d", x, got, expected)
		}
	}
}
