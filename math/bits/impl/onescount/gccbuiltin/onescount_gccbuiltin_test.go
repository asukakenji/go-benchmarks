package gccbuiltin_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/gccbuiltin"
)

func TestOnesCountAllNaive(t *testing.T) {
	// onescount.NaiveTest(t, gccbuiltin.OnesCount)
	// onescount.NaiveTest8(t, gccbuiltin.OnesCount8)
	// onescount.NaiveTest16(t, gccbuiltin.OnesCount16)
	onescount.NaiveTest32(t, gccbuiltin.OnesCount32)
	onescount.NaiveTest64(t, gccbuiltin.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		// Uint:   gccbuiltin.OnesCount,
		// Uint8:  gccbuiltin.OnesCount8,
		// Uint16: gccbuiltin.OnesCount16,
		Uint32: gccbuiltin.OnesCount32,
		Uint64: gccbuiltin.OnesCount64,
	})
	// onescount.TableTest(t, gccbuiltin.OnesCount)
	// onescount.TableTest8(t, gccbuiltin.OnesCount8)
	// onescount.TableTest16(t, gccbuiltin.OnesCount16)
	onescount.TableTest32(t, gccbuiltin.OnesCount32)
	onescount.TableTest64(t, gccbuiltin.OnesCount64)
}
