package nlz3_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/leadingzeros/nlz3"
)

func TestLeadingZerosAllNaive(t *testing.T) {
	leadingzeros.NaiveTest(t, nlz3.LeadingZeros)
	leadingzeros.NaiveTest8(t, nlz3.LeadingZeros8)
	leadingzeros.NaiveTest16(t, nlz3.LeadingZeros16)
	leadingzeros.NaiveTest32(t, nlz3.LeadingZeros32)
	leadingzeros.NaiveTest64(t, nlz3.LeadingZeros64)
}

func TestLeadingZerosAllTable(t *testing.T) {
	leadingzeros.BasicTableTest(t, &leadingzeros.LeadingZerosFuncs{
		Uint:   nlz3.LeadingZeros,
		Uint8:  nlz3.LeadingZeros8,
		Uint16: nlz3.LeadingZeros16,
		Uint32: nlz3.LeadingZeros32,
		Uint64: nlz3.LeadingZeros64,
	})
	leadingzeros.TableTest(t, nlz3.LeadingZeros)
	leadingzeros.TableTest8(t, nlz3.LeadingZeros8)
	leadingzeros.TableTest16(t, nlz3.LeadingZeros16)
	leadingzeros.TableTest32(t, nlz3.LeadingZeros32)
	leadingzeros.TableTest64(t, nlz3.LeadingZeros64)
}
