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
				t.Errorf("OnesCount(%d) = %d, expected %d", x, got, expected)
			}
		}
	}
	if onesCountFuncs.Uint8 != nil {
		for x, expected := range pop8tab {
			got := onesCountFuncs.Uint8(uint8(x))
			if got != int(expected) {
				t.Errorf("OnesCount8(%d) = %d, expected %d", x, got, expected)
			}
		}
	}
	if onesCountFuncs.Uint16 != nil {
		for x, expected := range pop8tab {
			got := onesCountFuncs.Uint16(uint16(x))
			if got != int(expected) {
				t.Errorf("OnesCount16(%d) = %d, expected %d", x, got, expected)
			}
		}
	}
	if onesCountFuncs.Uint32 != nil {
		for x, expected := range pop8tab {
			got := onesCountFuncs.Uint32(uint32(x))
			if got != int(expected) {
				t.Errorf("OnesCount32(%d) = %d, expected %d", x, got, expected)
			}
		}
	}
	if onesCountFuncs.Uint64 != nil {
		for x, expected := range pop8tab {
			got := onesCountFuncs.Uint64(uint64(x))
			if got != int(expected) {
				t.Errorf("OnesCount64(%d) = %d, expected %d", x, got, expected)
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
			t.Errorf("OnesCount(%d) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest8(t *testing.T, onesCount8Func func(uint8) int) {
	for i := 0; i < 256; i++ {
		x := uint8(rand.Uint64())
		expected := table.OnesCount8(x)
		got := onesCount8Func(x)
		if got != expected {
			t.Errorf("OnesCount8(%d) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest16(t *testing.T, onesCount16Func func(uint16) int) {
	for i := 0; i < 256; i++ {
		x := uint16(rand.Uint64())
		expected := table.OnesCount16(x)
		got := onesCount16Func(x)
		if got != expected {
			t.Errorf("OnesCount16(%d) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest32(t *testing.T, onesCount32Func func(uint32) int) {
	for i := 0; i < 256; i++ {
		x := uint32(rand.Uint64())
		expected := table.OnesCount32(x)
		got := onesCount32Func(x)
		if got != expected {
			t.Errorf("OnesCount32(%d) = %d, expected %d", x, got, expected)
		}
	}
}

func TableTest64(t *testing.T, onesCount64Func func(uint64) int) {
	for i := 0; i < 256; i++ {
		x := rand.Uint64()
		expected := table.OnesCount64(x)
		got := onesCount64Func(x)
		if got != expected {
			t.Errorf("OnesCount64(%d) = %d, expected %d", x, got, expected)
		}
	}
}
