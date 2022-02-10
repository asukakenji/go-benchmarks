package tcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/naive"
)

func TestAgainstNaiveImplementation(t *testing.T, leadingZeros func(uint) int) {
	testAgainst(t, naive.LeadingZeros, leadingZeros)
}

func TestAgainstNaiveImplementation8(t *testing.T, leadingZeros func(uint8) int) {
	testAgainst8(t, naive.LeadingZeros8, leadingZeros)
}

func TestAgainstNaiveImplementation16(t *testing.T, leadingZeros func(uint16) int) {
	testAgainst16(t, naive.LeadingZeros16, leadingZeros)
}

func TestAgainstNaiveImplementation32(t *testing.T, leadingZeros func(uint32) int) {
	testAgainst32(t, naive.LeadingZeros32, leadingZeros)
}

func TestAgainstNaiveImplementation64(t *testing.T, leadingZeros func(uint64) int) {
	testAgainst64(t, naive.LeadingZeros64, leadingZeros)
}

func TestAgainstNaiveImplementationPtr(t *testing.T, leadingZeros func(uintptr) int) {
	testAgainstPtr(t, naive.LeadingZerosPtr, leadingZeros)
}
