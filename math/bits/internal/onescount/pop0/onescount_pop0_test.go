package pop0_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop0"
)

func TestOnesCountAllNaive(t *testing.T) {
	//onescount.NaiveTest(t, pop0.OnesCount)
	onescount.NaiveTest8(t, pop0.OnesCount8)
	onescount.NaiveTest16(t, pop0.OnesCount16)
	onescount.NaiveTest32(t, pop0.OnesCount32)
	onescount.NaiveTest64(t, pop0.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		Uint:   nil, // pop0.OnesCount
		Uint8:  pop0.OnesCount8,
		Uint16: pop0.OnesCount16,
		Uint32: pop0.OnesCount32,
		Uint64: pop0.OnesCount64,
	})
	//onescount.TableTest(t, pop0.OnesCount)
	onescount.TableTest8(t, pop0.OnesCount8)
	onescount.TableTest16(t, pop0.OnesCount16)
	onescount.TableTest32(t, pop0.OnesCount32)
	onescount.TableTest64(t, pop0.OnesCount64)
}
