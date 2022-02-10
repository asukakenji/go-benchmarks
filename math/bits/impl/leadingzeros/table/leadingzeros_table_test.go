package table_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/table"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/tcommon"
)

func TestLeadingZerosAllBasic(t *testing.T) {
	tcommon.TestBasic(t, table.LeadingZerosConcept)
	tcommon.TestBasic8(t, table.LeadingZeros8)
	tcommon.TestBasic16(t, table.LeadingZeros16)
	tcommon.TestBasic32(t, table.LeadingZeros32)
	tcommon.TestBasic64(t, table.LeadingZeros64)
	// tcommon.TestBasicPtr(t, table.LeadingZerosPtr)
}

func TestLeadingZerosAllCases(t *testing.T) {
	tcommon.TestLeadingZeros(t, table.LeadingZerosConcept)
	tcommon.TestLeadingZeros8(t, table.LeadingZeros8)
	tcommon.TestLeadingZeros16(t, table.LeadingZeros16)
	tcommon.TestLeadingZeros32(t, table.LeadingZeros32)
	tcommon.TestLeadingZeros64(t, table.LeadingZeros64)
	// tcommon.TestLeadingZerosPtr(t, table.LeadingZerosPtr)
}

func TestLeadingZerosAllNaive(t *testing.T) {
	tcommon.TestAgainstNaiveImplementation(t, table.LeadingZerosConcept)
	tcommon.TestAgainstNaiveImplementation8(t, table.LeadingZeros8)
	tcommon.TestAgainstNaiveImplementation16(t, table.LeadingZeros16)
	tcommon.TestAgainstNaiveImplementation32(t, table.LeadingZeros32)
	tcommon.TestAgainstNaiveImplementation64(t, table.LeadingZeros64)
	// tcommon.TestAgainstNaiveImplementationPtr(t, table.LeadingZerosPtr)
}
