package pop1_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1"
)

func TestOnesCountAllNaive(t *testing.T) {
	//onescount.NaiveTest(t, pop1.OnesCount)
	onescount.NaiveTest8(t, pop1.OnesCount8)
	onescount.NaiveTest16(t, pop1.OnesCount16)
	onescount.NaiveTest32(t, pop1.OnesCount32)
	onescount.NaiveTest64(t, pop1.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		Uint:   nil, // pop1.OnesCount
		Uint8:  pop1.OnesCount8,
		Uint16: pop1.OnesCount16,
		Uint32: pop1.OnesCount32,
		Uint64: pop1.OnesCount64,
	})
	//onescount.TableTest(t, pop1.OnesCount)
	onescount.TableTest8(t, pop1.OnesCount8)
	onescount.TableTest16(t, pop1.OnesCount16)
	onescount.TableTest32(t, pop1.OnesCount32)
	onescount.TableTest64(t, pop1.OnesCount64)
}
