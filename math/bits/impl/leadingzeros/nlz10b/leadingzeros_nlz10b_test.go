package nlz10b_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz10b"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/tcommon"
)

func TestLeadingZerosAllBasic(t *testing.T) {
	// tcommon.TestBasic(t, nlz10b.LeadingZeros)
	// tcommon.TestBasic8(t, nlz10b.LeadingZeros8)
	// tcommon.TestBasic16(t, nlz10b.LeadingZeros16)
	tcommon.TestBasic32(t, nlz10b.LeadingZeros32)
	tcommon.TestBasic32(t, nlz10b.LeadingZeros32NoMultiply)
	// tcommon.TestBasic64(t, nlz10b.LeadingZeros64)
	// tcommon.TestBasicPtr(t, nlz10b.LeadingZerosPtr)
}

func TestLeadingZerosAllNaive(t *testing.T) {
	// tcommon.TestAgainstNaiveImplementation(t, nlz10b.LeadingZeros)
	// tcommon.TestAgainstNaiveImplementation8(t, nlz10b.LeadingZeros8)
	// tcommon.TestAgainstNaiveImplementation16(t, nlz10b.LeadingZeros16)
	tcommon.TestAgainstNaiveImplementation32(t, nlz10b.LeadingZeros32)
	tcommon.TestAgainstNaiveImplementation32(t, nlz10b.LeadingZeros32NoMultiply)
	// tcommon.TestAgainstNaiveImplementation64(t, nlz10b.LeadingZeros64)
	// tcommon.TestAgainstNaiveImplementationPtr(t, nlz10b.LeadingZerosPtr)
}

func TestLeadingZerosAllTable(t *testing.T) {
	// tcommon.TestAgainstTableImplementation(t, nlz10b.LeadingZeros)
	// tcommon.TestAgainstTableImplementation8(t, nlz10b.LeadingZeros8)
	// tcommon.TestAgainstTableImplementation16(t, nlz10b.LeadingZeros16)
	tcommon.TestAgainstTableImplementation32(t, nlz10b.LeadingZeros32)
	tcommon.TestAgainstTableImplementation32(t, nlz10b.LeadingZeros32NoMultiply)
	// tcommon.TestAgainstTableImplementation64(t, nlz10b.LeadingZeros64)
	// tcommon.TestAgainstTableImplementationPtr(t, nlz10b.LeadingZerosPtr)
}
