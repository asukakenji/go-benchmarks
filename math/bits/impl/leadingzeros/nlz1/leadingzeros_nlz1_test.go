package nlz1_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz1"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/tcommon"
)

func TestLeadingZerosAllBasic(t *testing.T) {
	// tcommon.TestBasic(t, nlz1.LeadingZeros)
	tcommon.TestBasic8(t, nlz1.LeadingZeros8)
	tcommon.TestBasic16(t, nlz1.LeadingZeros16)
	tcommon.TestBasic32(t, nlz1.LeadingZeros32)
	tcommon.TestBasic64(t, nlz1.LeadingZeros64)
	// tcommon.TestBasicPtr(t, nlz1.LeadingZerosPtr)
}

func TestLeadingZerosAgainstStdlib(t *testing.T) {
	// tcommon.TestAgainstStdlibImplementation(t, nlz1.LeadingZeros)
	tcommon.TestAgainstStdlibImplementation8(t, nlz1.LeadingZeros8)
	tcommon.TestAgainstStdlibImplementation16(t, nlz1.LeadingZeros16)
	tcommon.TestAgainstStdlibImplementation32(t, nlz1.LeadingZeros32)
	tcommon.TestAgainstStdlibImplementation64(t, nlz1.LeadingZeros64)
	// tcommon.TestAgainstStdlibImplementationPtr(t, nlz1.LeadingZerosPtr)
}

func TestLeadingZerosAgainstNaive(t *testing.T) {
	// tcommon.TestAgainstNaiveImplementation(t, nlz1.LeadingZeros)
	tcommon.TestAgainstNaiveImplementation8(t, nlz1.LeadingZeros8)
	tcommon.TestAgainstNaiveImplementation16(t, nlz1.LeadingZeros16)
	tcommon.TestAgainstNaiveImplementation32(t, nlz1.LeadingZeros32)
	tcommon.TestAgainstNaiveImplementation64(t, nlz1.LeadingZeros64)
	// tcommon.TestAgainstNaiveImplementationPtr(t, nlz1.LeadingZerosPtr)
}

func TestLeadingZerosAgainstTable(t *testing.T) {
	// tcommon.TestAgainstTableImplementation(t, nlz1.LeadingZeros)
	tcommon.TestAgainstTableImplementation8(t, nlz1.LeadingZeros8)
	tcommon.TestAgainstTableImplementation16(t, nlz1.LeadingZeros16)
	tcommon.TestAgainstTableImplementation32(t, nlz1.LeadingZeros32)
	tcommon.TestAgainstTableImplementation64(t, nlz1.LeadingZeros64)
	// tcommon.TestAgainstTableImplementationPtr(t, nlz1.LeadingZerosPtr)
}
