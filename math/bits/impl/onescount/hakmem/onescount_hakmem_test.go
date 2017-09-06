package hakmem_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/hakmem"
)

func TestOnesCountAllNaive(t *testing.T) {
	// onescount.NaiveTest(t, hakmem.OnesCount)
	// onescount.NaiveTest8(t, hakmem.OnesCount8)
	// onescount.NaiveTest16(t, hakmem.OnesCount16)
	onescount.NaiveTest32(t, hakmem.OnesCount32)
	onescount.NaiveTest32(t, hakmem.OnesCount32Unrolled)
	onescount.NaiveTest64(t, hakmem.OnesCount64)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.BasicTableTest(t, &onescount.OnesCountFuncs{
		Uint:   nil, // hakmem.OnesCount,
		Uint8:  nil, // hakmem.OnesCount8,
		Uint16: nil, // hakmem.OnesCount16,
		Uint32: hakmem.OnesCount32,
		Uint64: hakmem.OnesCount64,
	})
	// onescount.TableTest(t, hakmem.OnesCount)
	// onescount.TableTest8(t, hakmem.OnesCount8)
	// onescount.TableTest16(t, hakmem.OnesCount16)
	onescount.TableTest32(t, hakmem.OnesCount32)
	onescount.TableTest32(t, hakmem.OnesCount32Unrolled)
	onescount.TableTest64(t, hakmem.OnesCount64)
}
