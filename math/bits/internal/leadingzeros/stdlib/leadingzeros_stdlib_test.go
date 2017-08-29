package stdlib_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/leadingzeros/stdlib"
)

func TestLeadingZerosAllNaive(t *testing.T) {
	leadingzeros.NaiveTest(t, stdlib.LeadingZeros)
	leadingzeros.NaiveTest8(t, stdlib.LeadingZeros8)
	leadingzeros.NaiveTest16(t, stdlib.LeadingZeros16)
	leadingzeros.NaiveTest32(t, stdlib.LeadingZeros32)
	leadingzeros.NaiveTest64(t, stdlib.LeadingZeros64)
}

func TestLeadingZerosAllTable(t *testing.T) {
	leadingzeros.BasicTableTest(t, &leadingzeros.LeadingZerosFuncs{
		Uint:   stdlib.LeadingZeros,
		Uint8:  stdlib.LeadingZeros8,
		Uint16: stdlib.LeadingZeros16,
		Uint32: stdlib.LeadingZeros32,
		Uint64: stdlib.LeadingZeros64,
	})
	leadingzeros.TableTest(t, stdlib.LeadingZeros)
	leadingzeros.TableTest8(t, stdlib.LeadingZeros8)
	leadingzeros.TableTest16(t, stdlib.LeadingZeros16)
	leadingzeros.TableTest32(t, stdlib.LeadingZeros32)
	leadingzeros.TableTest64(t, stdlib.LeadingZeros64)
}
