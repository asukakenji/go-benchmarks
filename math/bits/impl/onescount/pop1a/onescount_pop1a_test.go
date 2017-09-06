package pop1a_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
)

func TestOnesCountAllNaive(t *testing.T) {
	onescount.NaiveTest(t, pop1a.OnesCount)
	onescount.NaiveTest8(t, pop1a.OnesCount8)
	onescount.NaiveTest16(t, pop1a.OnesCount16)
	onescount.NaiveTest32(t, pop1a.OnesCount32)
	onescount.NaiveTest64(t, pop1a.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		Uint:   pop1a.OnesCount,
		Uint8:  pop1a.OnesCount8,
		Uint16: pop1a.OnesCount16,
		Uint32: pop1a.OnesCount32,
		Uint64: pop1a.OnesCount64,
	})
	onescount.TableTest(t, pop1a.OnesCount)
	onescount.TableTest8(t, pop1a.OnesCount8)
	onescount.TableTest16(t, pop1a.OnesCount16)
	onescount.TableTest32(t, pop1a.OnesCount32)
	onescount.TableTest64(t, pop1a.OnesCount64)
}
