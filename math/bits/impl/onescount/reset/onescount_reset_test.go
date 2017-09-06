package reset_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/reset"
)

func TestOnesCountAllNaive(t *testing.T) {
	onescount.NaiveTest(t, reset.OnesCount)
	onescount.NaiveTest8(t, reset.OnesCount8)
	onescount.NaiveTest16(t, reset.OnesCount16)
	onescount.NaiveTest32(t, reset.OnesCount32)
	onescount.NaiveTest64(t, reset.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		Uint:   reset.OnesCount,
		Uint8:  reset.OnesCount8,
		Uint16: reset.OnesCount16,
		Uint32: reset.OnesCount32,
		Uint64: reset.OnesCount64,
	})
	onescount.TableTest(t, reset.OnesCount)
	onescount.TableTest8(t, reset.OnesCount8)
	onescount.TableTest16(t, reset.OnesCount16)
	onescount.TableTest32(t, reset.OnesCount32)
	onescount.TableTest64(t, reset.OnesCount64)
}
