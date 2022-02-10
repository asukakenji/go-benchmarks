package nlz1a_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz1a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/tcommon"
)

func TestLeadingZerosAllBasic(t *testing.T) {
	// tcommon.TestBasic(t, nlz1a.LeadingZeros)
	tcommon.TestBasic8(t, nlz1a.LeadingZeros8)
	tcommon.TestBasic16(t, nlz1a.LeadingZeros16)
	tcommon.TestBasic32(t, nlz1a.LeadingZeros32)
	tcommon.TestBasic64(t, nlz1a.LeadingZeros64)
	// tcommon.TestBasicPtr(t, nlz1a.LeadingZerosPtr)
}

func TestLeadingZerosAgainstStdlib(t *testing.T) {
	// tcommon.TestAgainstStdlibImplementation(t, nlz1a.LeadingZeros)
	tcommon.TestAgainstStdlibImplementation8(t, nlz1a.LeadingZeros8)
	tcommon.TestAgainstStdlibImplementation16(t, nlz1a.LeadingZeros16)
	tcommon.TestAgainstStdlibImplementation32(t, nlz1a.LeadingZeros32)
	tcommon.TestAgainstStdlibImplementation64(t, nlz1a.LeadingZeros64)
	// tcommon.TestAgainstStdlibImplementationPtr(t, nlz1a.LeadingZerosPtr)
}

func TestLeadingZerosAgainstNaive(t *testing.T) {
	// tcommon.TestAgainstNaiveImplementation(t, nlz1a.LeadingZeros)
	tcommon.TestAgainstNaiveImplementation8(t, nlz1a.LeadingZeros8)
	tcommon.TestAgainstNaiveImplementation16(t, nlz1a.LeadingZeros16)
	tcommon.TestAgainstNaiveImplementation32(t, nlz1a.LeadingZeros32)
	tcommon.TestAgainstNaiveImplementation64(t, nlz1a.LeadingZeros64)
	// tcommon.TestAgainstNaiveImplementationPtr(t, nlz1a.LeadingZerosPtr)
}

func TestLeadingZerosAgainstTable(t *testing.T) {
	// tcommon.TestAgainstTableImplementation(t, nlz1a.LeadingZeros)
	tcommon.TestAgainstTableImplementation8(t, nlz1a.LeadingZeros8)
	tcommon.TestAgainstTableImplementation16(t, nlz1a.LeadingZeros16)
	tcommon.TestAgainstTableImplementation32(t, nlz1a.LeadingZeros32)
	tcommon.TestAgainstTableImplementation64(t, nlz1a.LeadingZeros64)
	// tcommon.TestAgainstTableImplementationPtr(t, nlz1a.LeadingZerosPtr)
}
