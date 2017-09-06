package asm_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/asm"
)

func TestOnesCountAllNaive(t *testing.T) {
	// onescount.NaiveTest(t, asm.OnesCount)
	// onescount.NaiveTest8(t, asm.OnesCount8)
	// onescount.NaiveTest16(t, asm.OnesCount16)
	// onescount.NaiveTest32(t, asm.OnesCount32)
	onescount.NaiveTest64(t, asm.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		// Uint:   asm.OnesCount,
		// Uint8:  asm.OnesCount8,
		// Uint16: asm.OnesCount16,
		// Uint32: asm.OnesCount32,
		Uint64: asm.OnesCount64,
	})
	// onescount.TableTest(t, asm.OnesCount)
	// onescount.TableTest8(t, asm.OnesCount8)
	// onescount.TableTest16(t, asm.OnesCount16)
	// onescount.TableTest32(t, asm.OnesCount32)
	onescount.TableTest64(t, asm.OnesCount64)
}
