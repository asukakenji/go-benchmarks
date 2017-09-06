package pop2a_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop2a"
)

func TestOnesCountAllNaive(t *testing.T) {
	// onescount.NaiveTest(t, pop2a.OnesCount)
	// onescount.NaiveTest8(t, pop2a.OnesCount8)
	// onescount.NaiveTest16(t, pop2a.OnesCount16)
	onescount.NaiveTest32(t, pop2a.OnesCount32)
	// onescount.NaiveTest64(t, pop2a.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		Uint:   nil, // pop2a.OnesCount,
		Uint8:  nil, // pop2a.OnesCount8,
		Uint16: nil, // pop2a.OnesCount16,
		Uint32: pop2a.OnesCount32,
		Uint64: nil, // pop2a.OnesCount64,
	})
	// onescount.TableTest(t, pop2a.OnesCount)
	// onescount.TableTest8(t, pop2a.OnesCount8)
	// onescount.TableTest16(t, pop2a.OnesCount16)
	onescount.TableTest32(t, pop2a.OnesCount32)
	// onescount.TableTest64(t, pop2a.OnesCount64)
}
