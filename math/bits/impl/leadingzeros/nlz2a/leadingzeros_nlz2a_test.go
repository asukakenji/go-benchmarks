package nlz2a_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz2a"
)

func TestLeadingZerosAllNaive(t *testing.T) {
	leadingzeros.NaiveTest(t, nlz2a.LeadingZeros)
	leadingzeros.NaiveTest8(t, nlz2a.LeadingZeros8)
	leadingzeros.NaiveTest16(t, nlz2a.LeadingZeros16)
	leadingzeros.NaiveTest32(t, nlz2a.LeadingZeros32)
	leadingzeros.NaiveTest64(t, nlz2a.LeadingZeros64)
}

func TestLeadingZerosAllTable(t *testing.T) {
	leadingzeros.BasicTableTest(t, &leadingzeros.LeadingZerosFuncs{
		Uint:   nlz2a.LeadingZeros,
		Uint8:  nlz2a.LeadingZeros8,
		Uint16: nlz2a.LeadingZeros16,
		Uint32: nlz2a.LeadingZeros32,
		Uint64: nlz2a.LeadingZeros64,
	})
	leadingzeros.TableTest(t, nlz2a.LeadingZeros)
	leadingzeros.TableTest8(t, nlz2a.LeadingZeros8)
	leadingzeros.TableTest16(t, nlz2a.LeadingZeros16)
	leadingzeros.TableTest32(t, nlz2a.LeadingZeros32)
	leadingzeros.TableTest64(t, nlz2a.LeadingZeros64)
}
