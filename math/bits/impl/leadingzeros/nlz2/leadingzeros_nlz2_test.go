package nlz2_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/nlz2"
)

func TestLeadingZerosAllNaive(t *testing.T) {
	// leadingzeros.NaiveTest(t, nlz2.LeadingZeros)
	leadingzeros.NaiveTest8(t, nlz2.LeadingZeros8)
	leadingzeros.NaiveTest16(t, nlz2.LeadingZeros16)
	leadingzeros.NaiveTest32(t, nlz2.LeadingZeros32)
	leadingzeros.NaiveTest64(t, nlz2.LeadingZeros64)
}

func TestLeadingZerosAllTable(t *testing.T) {
	leadingzeros.BasicTableTest(t, &leadingzeros.LeadingZerosFuncs{
		// Uint:   nlz2.LeadingZeros,
		Uint8:  nlz2.LeadingZeros8,
		Uint16: nlz2.LeadingZeros16,
		Uint32: nlz2.LeadingZeros32,
		Uint64: nlz2.LeadingZeros64,
	})
	// leadingzeros.TableTest(t, nlz2.LeadingZeros)
	leadingzeros.TableTest8(t, nlz2.LeadingZeros8)
	leadingzeros.TableTest16(t, nlz2.LeadingZeros16)
	leadingzeros.TableTest32(t, nlz2.LeadingZeros32)
	leadingzeros.TableTest64(t, nlz2.LeadingZeros64)
}
