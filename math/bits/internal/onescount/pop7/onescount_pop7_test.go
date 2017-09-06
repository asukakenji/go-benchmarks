package pop7_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop7"
)

func TestOnesCountAllNaive(t *testing.T) {
	// onescount.NaiveTest(t, pop7.OnesCount)
	onescount.NaiveTest8(t, pop7.OnesCount8)
	onescount.NaiveTest8(t, pop7.OnesCount8Unrolled)
	// onescount.NaiveTest16(t, pop7.OnesCount16)
	// onescount.NaiveTest32(t, pop7.OnesCount32)
	// onescount.NaiveTest64(t, pop7.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		Uint:   nil, // pop7.OnesCount,
		Uint8:  pop7.OnesCount8,
		Uint16: nil, // pop7.OnesCount16,
		Uint32: nil, // pop7.OnesCount32,
		Uint64: nil, // pop7.OnesCount64,
	})
	// onescount.TableTest(t, pop7.OnesCount)
	onescount.TableTest8(t, pop7.OnesCount8)
	onescount.TableTest8(t, pop7.OnesCount8Unrolled)
	// onescount.TableTest16(t, pop7.OnesCount16)
	// onescount.TableTest32(t, pop7.OnesCount32)
	// onescount.TableTest64(t, pop7.OnesCount64)
}
