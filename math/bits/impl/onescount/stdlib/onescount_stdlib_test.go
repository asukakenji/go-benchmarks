package stdlib_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/stdlib"
)

func TestOnesCountAllNaive(t *testing.T) {
	onescount.NaiveTest(t, stdlib.OnesCount)
	onescount.NaiveTest8(t, stdlib.OnesCount8)
	onescount.NaiveTest16(t, stdlib.OnesCount16)
	onescount.NaiveTest32(t, stdlib.OnesCount32)
	onescount.NaiveTest64(t, stdlib.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		Uint:   stdlib.OnesCount,
		Uint8:  stdlib.OnesCount8,
		Uint16: stdlib.OnesCount16,
		Uint32: stdlib.OnesCount32,
		Uint64: stdlib.OnesCount64,
	})
	onescount.TableTest(t, stdlib.OnesCount)
	onescount.TableTest8(t, stdlib.OnesCount8)
	onescount.TableTest16(t, stdlib.OnesCount16)
	onescount.TableTest32(t, stdlib.OnesCount32)
	onescount.TableTest64(t, stdlib.OnesCount64)
}
