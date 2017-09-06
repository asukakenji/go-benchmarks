package pop8_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop8"
)

func TestOnesCountAllNaive(t *testing.T) {
	// onescount.NaiveTest(t, pop8.OnesCount)
	onescount.NaiveTest7(t, pop8.OnesCount7)
	onescount.NaiveTest7(t, pop8.OnesCount7Unrolled)
	onescount.NaiveTest8(t, pop8.OnesCount8)
	onescount.NaiveTest8(t, pop8.OnesCount8Unrolled)
	// onescount.NaiveTest16(t, pop8.OnesCount16)
	// onescount.NaiveTest32(t, pop8.OnesCount32)
	// onescount.NaiveTest64(t, pop8.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		Uint:   nil, // pop8.OnesCount,
		Uint8:  pop8.OnesCount8,
		Uint16: nil, // pop8.OnesCount16,
		Uint32: nil, // pop8.OnesCount32,
		Uint64: nil, // pop8.OnesCount64,
	})
	// onescount.TableTest(t, pop8.OnesCount)
	onescount.TableTest7(t, pop8.OnesCount7)
	onescount.TableTest7(t, pop8.OnesCount7Unrolled)
	onescount.TableTest8(t, pop8.OnesCount8)
	onescount.TableTest8(t, pop8.OnesCount8Unrolled)
	// onescount.TableTest16(t, pop8.OnesCount16)
	// onescount.TableTest32(t, pop8.OnesCount32)
	// onescount.TableTest64(t, pop8.OnesCount64)
}
