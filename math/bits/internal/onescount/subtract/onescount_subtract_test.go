package subtract_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/subtract"
)

func TestOnesCountAllNaive(t *testing.T) {
	onescount.NaiveTest(t, subtract.OnesCount)
	onescount.NaiveTest8(t, subtract.OnesCount8)
	onescount.NaiveTest16(t, subtract.OnesCount16)
	onescount.NaiveTest32(t, subtract.OnesCount32)
	onescount.NaiveTest64(t, subtract.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		Uint:   subtract.OnesCount,
		Uint8:  subtract.OnesCount8,
		Uint16: subtract.OnesCount16,
		Uint32: subtract.OnesCount32,
		Uint64: subtract.OnesCount64,
	})
	onescount.TableTest(t, subtract.OnesCount)
	onescount.TableTest8(t, subtract.OnesCount8)
	onescount.TableTest16(t, subtract.OnesCount16)
	onescount.TableTest32(t, subtract.OnesCount32)
	onescount.TableTest64(t, subtract.OnesCount64)
}
