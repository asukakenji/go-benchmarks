package tcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/naive"
)

func TestAgainstNaiveImplementation(t *testing.T, leadingZeros func(uint) int) {
	testAgainst(t, naive.LeadingZeros, leadingZeros)
}

func TestAgainstNaiveImplementation8(t *testing.T, leadingZeros8 func(uint8) int) {
	testAgainst8(t, naive.LeadingZeros8, leadingZeros8)
}

func TestAgainstNaiveImplementation16(t *testing.T, leadingZeros16 func(uint16) int) {
	testAgainst16(t, naive.LeadingZeros16, leadingZeros16)
}

func TestAgainstNaiveImplementation32(t *testing.T, leadingZeros32 func(uint32) int) {
	testAgainst32(t, naive.LeadingZeros32, leadingZeros32)
}

func TestAgainstNaiveImplementation64(t *testing.T, leadingZeros64 func(uint64) int) {
	testAgainst64(t, naive.LeadingZeros64, leadingZeros64)
}

func TestAgainstNaiveImplementationPtr(t *testing.T, leadingZerosPtr func(uintptr) int) {
	testAgainstPtr(t, naive.LeadingZerosPtr, leadingZerosPtr)
}
