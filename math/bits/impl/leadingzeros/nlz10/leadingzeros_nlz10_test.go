package nlz10_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz10"
)

func TestLeadingZerosAllNaive(t *testing.T) {
	// leadingzeros.NaiveTest(t, nlz10.LeadingZeros)
	// leadingzeros.NaiveTest8(t, nlz10.LeadingZeros8)
	// leadingzeros.NaiveTest16(t, nlz10.LeadingZeros16)
	leadingzeros.NaiveTest32(t, nlz10.LeadingZeros32)
	leadingzeros.NaiveTest32(t, nlz10.LeadingZeros32NoMultiply)
	// leadingzeros.NaiveTest64(t, nlz10.LeadingZeros64)
}

func TestLeadingZerosAllTable(t *testing.T) {
	leadingzeros.BasicTableTest(t, &leadingzeros.LeadingZerosFuncs{
		// Uint:   nlz10.LeadingZeros,
		// Uint8:  nlz10.LeadingZeros8,
		// Uint16: nlz10.LeadingZeros16,
		Uint32: nlz10.LeadingZeros32,
		// Uint64: nlz10.LeadingZeros64,
	})
	// leadingzeros.TableTest(t, nlz10.LeadingZeros)
	// leadingzeros.TableTest8(t, nlz10.LeadingZeros8)
	// leadingzeros.TableTest16(t, nlz10.LeadingZeros16)
	leadingzeros.TableTest32(t, nlz10.LeadingZeros32)
	leadingzeros.TableTest32(t, nlz10.LeadingZeros32NoMultiply)
	// leadingzeros.TableTest64(t, nlz10.LeadingZeros64)
}
