package naive_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/naive"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/tcommon"
)

func TestLeadingZerosAllBasic(t *testing.T) {
	tcommon.TestBasic(t, naive.LeadingZeros)
	tcommon.TestBasic8(t, naive.LeadingZeros8)
	tcommon.TestBasic16(t, naive.LeadingZeros16)
	tcommon.TestBasic32(t, naive.LeadingZeros32)
	tcommon.TestBasic64(t, naive.LeadingZeros64)
	tcommon.TestBasicPtr(t, naive.LeadingZerosPtr)
	// ---
	tcommon.TestBasic(t, naive.LeadingZerosGeneric[uint])
	tcommon.TestBasic8(t, naive.LeadingZerosGeneric[uint8])
	tcommon.TestBasic16(t, naive.LeadingZerosGeneric[uint16])
	tcommon.TestBasic32(t, naive.LeadingZerosGeneric[uint32])
	tcommon.TestBasic64(t, naive.LeadingZerosGeneric[uint64])
	tcommon.TestBasicPtr(t, naive.LeadingZerosGeneric[uintptr])
}

func TestLeadingZerosAllCases(t *testing.T) {
	tcommon.TestLeadingZeros(t, naive.LeadingZeros)
	tcommon.TestLeadingZeros8(t, naive.LeadingZeros8)
	tcommon.TestLeadingZeros16(t, naive.LeadingZeros16)
	tcommon.TestLeadingZeros32(t, naive.LeadingZeros32)
	tcommon.TestLeadingZeros64(t, naive.LeadingZeros64)
	tcommon.TestLeadingZerosPtr(t, naive.LeadingZerosPtr)
	// ---
	tcommon.TestLeadingZeros(t, naive.LeadingZerosGeneric[uint])
	tcommon.TestLeadingZeros8(t, naive.LeadingZerosGeneric[uint8])
	tcommon.TestLeadingZeros16(t, naive.LeadingZerosGeneric[uint16])
	tcommon.TestLeadingZeros32(t, naive.LeadingZerosGeneric[uint32])
	tcommon.TestLeadingZeros64(t, naive.LeadingZerosGeneric[uint64])
	tcommon.TestLeadingZerosPtr(t, naive.LeadingZerosGeneric[uintptr])
}

func TestLeadingZerosAllTable(t *testing.T) {
	tcommon.TestAgainstTableImplementation(t, naive.LeadingZeros)
	tcommon.TestAgainstTableImplementation8(t, naive.LeadingZeros8)
	tcommon.TestAgainstTableImplementation16(t, naive.LeadingZeros16)
	tcommon.TestAgainstTableImplementation32(t, naive.LeadingZeros32)
	tcommon.TestAgainstTableImplementation64(t, naive.LeadingZeros64)
	tcommon.TestAgainstTableImplementationPtr(t, naive.LeadingZerosPtr)
	// ---
	tcommon.TestAgainstTableImplementation(t, naive.LeadingZerosGeneric[uint])
	tcommon.TestAgainstTableImplementation8(t, naive.LeadingZerosGeneric[uint8])
	tcommon.TestAgainstTableImplementation16(t, naive.LeadingZerosGeneric[uint16])
	tcommon.TestAgainstTableImplementation32(t, naive.LeadingZerosGeneric[uint32])
	tcommon.TestAgainstTableImplementation64(t, naive.LeadingZerosGeneric[uint64])
	tcommon.TestAgainstTableImplementationPtr(t, naive.LeadingZerosGeneric[uintptr])
}
