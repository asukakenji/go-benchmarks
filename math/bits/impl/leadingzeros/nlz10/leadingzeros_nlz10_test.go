package nlz10_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz10"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/tcommon"
)

func TestLeadingZerosAllBasic(t *testing.T) {
	// tcommon.TestBasic(t, nlz10.LeadingZeros)
	// tcommon.TestBasic8(t, nlz10.LeadingZeros8)
	// tcommon.TestBasic16(t, nlz10.LeadingZeros16)
	tcommon.TestBasic32(t, nlz10.LeadingZeros32)
	tcommon.TestBasic32(t, nlz10.LeadingZeros32NoMultiply)
	// tcommon.TestBasic64(t, nlz10.LeadingZeros64)
	// tcommon.TestBasicPtr(t, nlz10.LeadingZerosPtr)
}

func TestLeadingZerosAllNaive(t *testing.T) {
	// tcommon.TestAgainstNaiveImplementation(t, nlz10.LeadingZeros)
	// tcommon.TestAgainstNaiveImplementation8(t, nlz10.LeadingZeros8)
	// tcommon.TestAgainstNaiveImplementation16(t, nlz10.LeadingZeros16)
	tcommon.TestAgainstNaiveImplementation32(t, nlz10.LeadingZeros32)
	tcommon.TestAgainstNaiveImplementation32(t, nlz10.LeadingZeros32NoMultiply)
	// tcommon.TestAgainstNaiveImplementation64(t, nlz10.LeadingZeros64)
	// tcommon.TestAgainstNaiveImplementationPtr(t, nlz10.LeadingZerosPtr)
}

func TestLeadingZerosAllTable(t *testing.T) {
	// tcommon.TestAgainstTableImplementation(t, nlz10.LeadingZeros)
	// tcommon.TestAgainstTableImplementation8(t, nlz10.LeadingZeros8)
	// tcommon.TestAgainstTableImplementation16(t, nlz10.LeadingZeros16)
	tcommon.TestAgainstTableImplementation32(t, nlz10.LeadingZeros32)
	tcommon.TestAgainstTableImplementation32(t, nlz10.LeadingZeros32NoMultiply)
	// tcommon.TestAgainstTableImplementation64(t, nlz10.LeadingZeros64)
	// tcommon.TestAgainstTableImplementationPtr(t, nlz10.LeadingZerosPtr)
}
