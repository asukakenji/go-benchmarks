package nlz5_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz5"
)

func TestLeadingZerosAllNaive(t *testing.T) {
	// leadingzeros.NaiveTest(t, nlz5.LeadingZeros)
	leadingzeros.NaiveTest8(t, nlz5.LeadingZeros8)
	leadingzeros.NaiveTest16(t, nlz5.LeadingZeros16)
	leadingzeros.NaiveTest32(t, nlz5.LeadingZeros32)
	leadingzeros.NaiveTest64(t, nlz5.LeadingZeros64)
}

func TestLeadingZerosAllTable(t *testing.T) {
	leadingzeros.BasicTableTest(t, &leadingzeros.LeadingZerosFuncs{
		// Uint:   nlz5.LeadingZeros,
		Uint8:  nlz5.LeadingZeros8,
		Uint16: nlz5.LeadingZeros16,
		Uint32: nlz5.LeadingZeros32,
		Uint64: nlz5.LeadingZeros64,
	})
	// leadingzeros.TableTest(t, nlz5.LeadingZeros)
	leadingzeros.TableTest8(t, nlz5.LeadingZeros8)
	leadingzeros.TableTest16(t, nlz5.LeadingZeros16)
	leadingzeros.TableTest32(t, nlz5.LeadingZeros32)
	leadingzeros.TableTest64(t, nlz5.LeadingZeros64)
}
