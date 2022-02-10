package stdlib_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/stdlib"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/tcommon"
)

func TestLeadingZerosAllBasic(t *testing.T) {
	tcommon.TestBasic(t, stdlib.LeadingZeros)
	tcommon.TestBasic8(t, stdlib.LeadingZeros8)
	tcommon.TestBasic16(t, stdlib.LeadingZeros16)
	tcommon.TestBasic32(t, stdlib.LeadingZeros32)
	tcommon.TestBasic64(t, stdlib.LeadingZeros64)
	tcommon.TestBasicPtr(t, stdlib.LeadingZerosPtr)
}

func TestLeadingZerosAllCases(t *testing.T) {
	tcommon.TestLeadingZeros(t, stdlib.LeadingZeros)
	tcommon.TestLeadingZeros8(t, stdlib.LeadingZeros8)
	tcommon.TestLeadingZeros16(t, stdlib.LeadingZeros16)
	tcommon.TestLeadingZeros32(t, stdlib.LeadingZeros32)
	tcommon.TestLeadingZeros64(t, stdlib.LeadingZeros64)
	tcommon.TestLeadingZerosPtr(t, stdlib.LeadingZerosPtr)
}

func TestLeadingZerosAllTable(t *testing.T) {
	tcommon.TestAgainstTableImplementation(t, stdlib.LeadingZeros)
	tcommon.TestAgainstTableImplementation8(t, stdlib.LeadingZeros8)
	tcommon.TestAgainstTableImplementation16(t, stdlib.LeadingZeros16)
	tcommon.TestAgainstTableImplementation32(t, stdlib.LeadingZeros32)
	tcommon.TestAgainstTableImplementation64(t, stdlib.LeadingZeros64)
	tcommon.TestAgainstTableImplementationPtr(t, stdlib.LeadingZerosPtr)
}
