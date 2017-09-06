package pop3_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop3"
)

func TestOnesCountAllNaive(t *testing.T) {
	// onescount.NaiveTest(t, pop3.OnesCount)
	// onescount.NaiveTest8(t, pop3.OnesCount8)
	// onescount.NaiveTest16(t, pop3.OnesCount16)
	onescount.NaiveTest32(t, pop3.OnesCount32)
	onescount.NaiveTest64(t, pop3.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		// Uint:   pop3.OnesCount,
		// Uint8:  pop3.OnesCount8,
		// Uint16: pop3.OnesCount16,
		Uint32: pop3.OnesCount32,
		Uint64: pop3.OnesCount64,
	})
	// onescount.TableTest(t, pop3.OnesCount)
	// onescount.TableTest8(t, pop3.OnesCount8)
	// onescount.TableTest16(t, pop3.OnesCount16)
	onescount.TableTest32(t, pop3.OnesCount32)
	onescount.TableTest64(t, pop3.OnesCount64)
}
