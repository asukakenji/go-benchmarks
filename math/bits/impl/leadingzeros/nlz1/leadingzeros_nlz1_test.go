package nlz1_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz1"
)

func TestLeadingZerosAllNaive(t *testing.T) {
	// leadingzeros.NaiveTest(t, nlz1.LeadingZeros)
	leadingzeros.NaiveTest8(t, nlz1.LeadingZeros8)
	leadingzeros.NaiveTest16(t, nlz1.LeadingZeros16)
	leadingzeros.NaiveTest32(t, nlz1.LeadingZeros32)
	leadingzeros.NaiveTest64(t, nlz1.LeadingZeros64)
}

func TestLeadingZerosAllTable(t *testing.T) {
	leadingzeros.BasicTableTest(t, &leadingzeros.LeadingZerosFuncs{
		// Uint:   nlz1.LeadingZeros,
		Uint8:  nlz1.LeadingZeros8,
		Uint16: nlz1.LeadingZeros16,
		Uint32: nlz1.LeadingZeros32,
		Uint64: nlz1.LeadingZeros64,
	})
	// leadingzeros.TableTest(t, nlz1.LeadingZeros)
	leadingzeros.TableTest8(t, nlz1.LeadingZeros8)
	leadingzeros.TableTest16(t, nlz1.LeadingZeros16)
	leadingzeros.TableTest32(t, nlz1.LeadingZeros32)
	leadingzeros.TableTest64(t, nlz1.LeadingZeros64)
}
