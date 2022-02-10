package nlz2_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/tcommon"
)

func TestLeadingZerosAllBasic(t *testing.T) {
	// tcommon.TestBasic(t, nlz2.LeadingZeros)
	tcommon.TestBasic8(t, nlz2.LeadingZeros8)
	tcommon.TestBasic16(t, nlz2.LeadingZeros16)
	tcommon.TestBasic32(t, nlz2.LeadingZeros32)
	tcommon.TestBasic64(t, nlz2.LeadingZeros64)
	// tcommon.TestBasicPtr(t, nlz2.LeadingZerosPtr)
}

func TestLeadingZerosAgainstStdlib(t *testing.T) {
	// tcommon.TestAgainstStdlibImplementation(t, nlz2.LeadingZeros)
	tcommon.TestAgainstStdlibImplementation8(t, nlz2.LeadingZeros8)
	tcommon.TestAgainstStdlibImplementation16(t, nlz2.LeadingZeros16)
	tcommon.TestAgainstStdlibImplementation32(t, nlz2.LeadingZeros32)
	tcommon.TestAgainstStdlibImplementation64(t, nlz2.LeadingZeros64)
	// tcommon.TestAgainstStdlibImplementationPtr(t, nlz2.LeadingZerosPtr)
}

func TestLeadingZerosAgainstNaive(t *testing.T) {
	// tcommon.TestAgainstNaiveImplementation(t, nlz2.LeadingZeros)
	tcommon.TestAgainstNaiveImplementation8(t, nlz2.LeadingZeros8)
	tcommon.TestAgainstNaiveImplementation16(t, nlz2.LeadingZeros16)
	tcommon.TestAgainstNaiveImplementation32(t, nlz2.LeadingZeros32)
	tcommon.TestAgainstNaiveImplementation64(t, nlz2.LeadingZeros64)
	// tcommon.TestAgainstNaiveImplementationPtr(t, nlz2.LeadingZerosPtr)
}

func TestLeadingZerosAgainstTable(t *testing.T) {
	// tcommon.TestAgainstTableImplementation(t, nlz2.LeadingZeros)
	tcommon.TestAgainstTableImplementation8(t, nlz2.LeadingZeros8)
	tcommon.TestAgainstTableImplementation16(t, nlz2.LeadingZeros16)
	tcommon.TestAgainstTableImplementation32(t, nlz2.LeadingZeros32)
	tcommon.TestAgainstTableImplementation64(t, nlz2.LeadingZeros64)
	// tcommon.TestAgainstTableImplementationPtr(t, nlz2.LeadingZerosPtr)
}
