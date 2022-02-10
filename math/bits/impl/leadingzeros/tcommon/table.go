package tcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/table"
)

func TestAgainstTableImplementation(t *testing.T, leadingZeros func(uint) int) {
	testAgainst(t, table.LeadingZerosConcept, leadingZeros)
}

func TestAgainstTableImplementation8(t *testing.T, leadingZeros func(uint8) int) {
	testAgainst8(t, table.LeadingZeros8, leadingZeros)
}

func TestAgainstTableImplementation16(t *testing.T, leadingZeros func(uint16) int) {
	testAgainst16(t, table.LeadingZeros16, leadingZeros)
}

func TestAgainstTableImplementation32(t *testing.T, leadingZeros func(uint32) int) {
	testAgainst32(t, table.LeadingZeros32, leadingZeros)
}

func TestAgainstTableImplementation64(t *testing.T, leadingZeros func(uint64) int) {
	testAgainst64(t, table.LeadingZeros64, leadingZeros)
}

func TestAgainstTableImplementationPtr(t *testing.T, leadingZeros func(uintptr) int) {
	// TODO: Write this!
	// testAgainstPtr(t, table.LeadingZerosPtr, leadingZeros)
}
