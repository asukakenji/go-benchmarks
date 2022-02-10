package nlz2a_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2a"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/tcommon"
)

func TestLeadingZerosAllBasic(t *testing.T) {
	tcommon.TestBasic(t, nlz2a.LeadingZeros)
	tcommon.TestBasic8(t, nlz2a.LeadingZeros8)
	tcommon.TestBasic16(t, nlz2a.LeadingZeros16)
	tcommon.TestBasic32(t, nlz2a.LeadingZeros32)
	tcommon.TestBasic64(t, nlz2a.LeadingZeros64)
	tcommon.TestBasicPtr(t, nlz2a.LeadingZerosPtr)
}

func TestLeadingZerosAllNaive(t *testing.T) {
	tcommon.TestAgainstNaiveImplementation(t, nlz2a.LeadingZeros)
	tcommon.TestAgainstNaiveImplementation8(t, nlz2a.LeadingZeros8)
	tcommon.TestAgainstNaiveImplementation16(t, nlz2a.LeadingZeros16)
	tcommon.TestAgainstNaiveImplementation32(t, nlz2a.LeadingZeros32)
	tcommon.TestAgainstNaiveImplementation64(t, nlz2a.LeadingZeros64)
	tcommon.TestAgainstNaiveImplementationPtr(t, nlz2a.LeadingZerosPtr)
}

func TestLeadingZerosAllTable(t *testing.T) {
	tcommon.TestAgainstTableImplementation(t, nlz2a.LeadingZeros)
	tcommon.TestAgainstTableImplementation8(t, nlz2a.LeadingZeros8)
	tcommon.TestAgainstTableImplementation16(t, nlz2a.LeadingZeros16)
	tcommon.TestAgainstTableImplementation32(t, nlz2a.LeadingZeros32)
	tcommon.TestAgainstTableImplementation64(t, nlz2a.LeadingZeros64)
	tcommon.TestAgainstTableImplementationPtr(t, nlz2a.LeadingZerosPtr)
}
