package tcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/table"
)

func TestAgainstTableImplementation(t *testing.T, leadingZeros func(uint) int) {
	testAgainst(t, table.LeadingZeros, leadingZeros)
}

func TestAgainstTableImplementation8(t *testing.T, leadingZeros8 func(uint8) int) {
	testAgainst8(t, table.LeadingZeros8, leadingZeros8)
}

func TestAgainstTableImplementation16(t *testing.T, leadingZeros16 func(uint16) int) {
	testAgainst16(t, table.LeadingZeros16, leadingZeros16)
}

func TestAgainstTableImplementation32(t *testing.T, leadingZeros32 func(uint32) int) {
	testAgainst32(t, table.LeadingZeros32, leadingZeros32)
}

func TestAgainstTableImplementation64(t *testing.T, leadingZeros64 func(uint64) int) {
	testAgainst64(t, table.LeadingZeros64, leadingZeros64)
}

func TestAgainstTableImplementationPtr(t *testing.T, leadingZerosPtr func(uintptr) int) {
	testAgainstPtr(t, table.LeadingZerosPtr, leadingZerosPtr)
}
