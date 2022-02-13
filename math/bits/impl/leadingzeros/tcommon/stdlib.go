package tcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/stdlib"
)

func TestAgainstStdlibImplementation(t *testing.T, leadingZeros func(uint) int) {
	testAgainst(t, stdlib.LeadingZeros, leadingZeros)
}

func TestAgainstStdlibImplementation8(t *testing.T, leadingZeros8 func(uint8) int) {
	testAgainst8(t, stdlib.LeadingZeros8, leadingZeros8)
}

func TestAgainstStdlibImplementation16(t *testing.T, leadingZeros16 func(uint16) int) {
	testAgainst16(t, stdlib.LeadingZeros16, leadingZeros16)
}

func TestAgainstStdlibImplementation32(t *testing.T, leadingZeros32 func(uint32) int) {
	testAgainst32(t, stdlib.LeadingZeros32, leadingZeros32)
}

func TestAgainstStdlibImplementation64(t *testing.T, leadingZeros64 func(uint64) int) {
	testAgainst64(t, stdlib.LeadingZeros64, leadingZeros64)
}

func TestAgainstStdlibImplementationPtr(t *testing.T, leadingZerosPtr func(uintptr) int) {
	testAgainstPtr(t, stdlib.LeadingZerosPtr, leadingZerosPtr)
}
