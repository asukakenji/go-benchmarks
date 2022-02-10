package table_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/table"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/tcommon"
)

func TestLeadingZerosAllBasic(t *testing.T) {
	tcommon.TestBasic(t, table.LeadingZeros)
	tcommon.TestBasic8(t, table.LeadingZeros8)
	tcommon.TestBasic16(t, table.LeadingZeros16)
	tcommon.TestBasic32(t, table.LeadingZeros32)
	tcommon.TestBasic64(t, table.LeadingZeros64)
	tcommon.TestBasicPtr(t, table.LeadingZerosPtr)
}

func TestLeadingZerosAllCases(t *testing.T) {
	tcommon.TestLeadingZeros(t, table.LeadingZeros)
	tcommon.TestLeadingZeros8(t, table.LeadingZeros8)
	tcommon.TestLeadingZeros16(t, table.LeadingZeros16)
	tcommon.TestLeadingZeros32(t, table.LeadingZeros32)
	tcommon.TestLeadingZeros64(t, table.LeadingZeros64)
	tcommon.TestLeadingZerosPtr(t, table.LeadingZerosPtr)
}

func TestLeadingZerosAgainstStdlib(t *testing.T) {
	tcommon.TestAgainstStdlibImplementation(t, table.LeadingZeros)
	tcommon.TestAgainstStdlibImplementation8(t, table.LeadingZeros8)
	tcommon.TestAgainstStdlibImplementation16(t, table.LeadingZeros16)
	tcommon.TestAgainstStdlibImplementation32(t, table.LeadingZeros32)
	tcommon.TestAgainstStdlibImplementation64(t, table.LeadingZeros64)
	tcommon.TestAgainstStdlibImplementationPtr(t, table.LeadingZerosPtr)
}

func TestLeadingZerosAgainstNaive(t *testing.T) {
	tcommon.TestAgainstNaiveImplementation(t, table.LeadingZeros)
	tcommon.TestAgainstNaiveImplementation8(t, table.LeadingZeros8)
	tcommon.TestAgainstNaiveImplementation16(t, table.LeadingZeros16)
	tcommon.TestAgainstNaiveImplementation32(t, table.LeadingZeros32)
	tcommon.TestAgainstNaiveImplementation64(t, table.LeadingZeros64)
	tcommon.TestAgainstNaiveImplementationPtr(t, table.LeadingZerosPtr)
}
