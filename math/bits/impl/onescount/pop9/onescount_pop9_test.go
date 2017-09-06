package pop9_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop9"
)

func TestOnesCountAllNaive(t *testing.T) {
	// onescount.NaiveTest(t, pop9.OnesCount)
	onescount.NaiveTest8(t, pop9.OnesCount8)
	onescount.NaiveTest8(t, pop9.OnesCount8Unrolled)
	onescount.NaiveTest15(t, pop9.OnesCount15)
	onescount.NaiveTest15(t, pop9.OnesCount15Unrolled)
	onescount.NaiveTest16(t, pop9.OnesCount16)
	onescount.NaiveTest16(t, pop9.OnesCount16Unrolled)
	// onescount.NaiveTest32(t, pop9.OnesCount32)
	// onescount.NaiveTest64(t, pop9.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		Uint:   nil, // pop9.OnesCount,
		Uint8:  pop9.OnesCount8,
		Uint16: pop9.OnesCount16,
		Uint32: nil, // pop9.OnesCount32,
		Uint64: nil, // pop9.OnesCount64,
	})
	// onescount.TableTest(t, pop9.OnesCount)
	onescount.TableTest8(t, pop9.OnesCount8)
	onescount.TableTest8(t, pop9.OnesCount8Unrolled)
	onescount.TableTest15(t, pop9.OnesCount15)
	onescount.TableTest15(t, pop9.OnesCount15Unrolled)
	onescount.TableTest16(t, pop9.OnesCount16)
	onescount.TableTest16(t, pop9.OnesCount16Unrolled)
	// onescount.TableTest32(t, pop9.OnesCount32)
	// onescount.TableTest64(t, pop9.OnesCount64)
}
